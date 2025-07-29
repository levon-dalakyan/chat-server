package chats

import (
	"github.com/ldevprog/chat-server/internal/service"
	desc "github.com/ldevprog/chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatsService service.ChatsService
}

func NewImplementation(chatsService service.ChatsService) *Implementation {
	return &Implementation{
		chatsService: chatsService,
	}
}
