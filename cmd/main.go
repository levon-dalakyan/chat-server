package main

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/levon-dalakyan/chat-server/pkg/chat_v1"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedChatV1Server
	db *pgxpool.Pool
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("id", "usernames").
		Values(randInt64Positive(), req.GetUsernames()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to build SQL query: %v", err)
	}

	var chatId int64
	err = s.db.QueryRow(ctx, query, args...).Scan(&chatId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert chat: %v", err)
	}

	return &desc.CreateResponse{
		Id: chatId,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	builderDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to build SQL query: %v", err)
	}

	res, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to delete chat: %v", err)
	}

	log.Printf("deleted %d rows", res.RowsAffected())

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	builderInsert := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("id", "chat_id", "sender", "text", "created_at").
		Values(randInt64Positive(), req.GetChatId(), req.GetFrom(), req.GetText(), req.GetTimestamp().AsTime())

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to build SQL query: %v", err)
	}

	res, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to insert message: %v", err)
	}

	log.Printf("inserted %d rows", res.RowsAffected())

	return &emptypb.Empty{}, nil
}

func randInt64Positive() int64 {
	var b [8]byte
	rand.Read(b[:])
	u := int64(binary.LittleEndian.Uint64(b[:]))
	return int64(u & 0x7FFFFFFFFFFFFFFF)
}

func getDSN() string {
	port := os.Getenv("PG_PORT")
	dbname := os.Getenv("PG_DATABASE_NAME")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASSWORD")

	return fmt.Sprintf(
		"host=localhost port=%s dbname=%s user=%s password=%s sslmode=disable",
		port,
		dbname,
		user,
		pass,
	)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
}

func main() {
	ctx := context.Background()
	dbDSN := getDSN()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{db: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
