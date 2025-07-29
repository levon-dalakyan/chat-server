package repository

import (
	"context"

	"github.com/ldevprog/chat-server/internal/model"
)

type ChatsRepository interface {
	Create(ctx context.Context, chat *model.ChatCreate) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message *model.MessageCreate) error
}
