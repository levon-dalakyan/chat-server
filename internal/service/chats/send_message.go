package chats

import (
	"context"

	"github.com/levon-dalakyan/chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, message *model.MessageCreate) error {
	err := s.chatsRepository.SendMessage(ctx, message)

	return err
}
