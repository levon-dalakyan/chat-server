package chats

import (
	"context"

	"github.com/ldevprog/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, chat *model.ChatCreate) (int64, error) {
	chatId, err := s.chatsRepository.Create(ctx, chat)
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
