package chats

import (
	"context"

	"github.com/ldevprog/chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, message *model.MessageCreate) error {
	err := s.chatsRepository.SendMessage(ctx, message)

	return err
}
