package chats

import (
	"github.com/levon-dalakyan/chat-server/internal/repository"
	"github.com/levon-dalakyan/chat-server/internal/service"
)

type serv struct {
	chatsRepository repository.ChatsRepository
}

func NewService(chatsRepository repository.ChatsRepository) service.ChatsService {
	return &serv{
		chatsRepository: chatsRepository,
	}
}
