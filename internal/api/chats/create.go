package chats

import (
	"context"

	"github.com/levon-dalakyan/chat-server/internal/converter"
	desc "github.com/levon-dalakyan/chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chatId, err := i.chatsService.Create(ctx, converter.ToChatCreateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: chatId,
	}, nil
}
