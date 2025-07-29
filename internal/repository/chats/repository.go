package chats

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ldevprog/platform-common/pkg/db"

	"github.com/ldevprog/chat-server/internal/helpers"
	"github.com/ldevprog/chat-server/internal/model"
	"github.com/ldevprog/chat-server/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatsRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, chat *model.ChatCreate) (int64, error) {
	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("id", "usernames").
		Values(helpers.RandInt64Positive(), chat.UserNames).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, status.Errorf(codes.Internal, "failed to build SQL query: %v", err)
	}

	q := db.Query{
		Name:     "chats_repository.Create",
		QueryRaw: query,
	}

	var chatId int64
	err = r.db.DB().ScanOneContext(ctx, &chatId, q, args...)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "failed to insert chat: %v", err)
	}

	return chatId, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to build SQL query: %v", err)
	}

	q := db.Query{
		Name:     "chats_repository.Delete",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to delete chat: %v", err)
	}

	log.Printf("deleted %d rows", res.RowsAffected())

	return nil
}

func (r *repo) SendMessage(ctx context.Context, message *model.MessageCreate) error {
	builderInsert := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("id", "chat_id", "sender", "text", "created_at").
		Values(
			helpers.RandInt64Positive(),
			message.ChatId,
			message.From,
			message.Text,
			message.Timestamp,
		)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to build SQL query: %v", err)
	}

	q := db.Query{
		Name:     "chats_repository.SendMessage",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to insert message: %v", err)
	}

	log.Printf("inserted %d rows", res.RowsAffected())

	return nil
}
