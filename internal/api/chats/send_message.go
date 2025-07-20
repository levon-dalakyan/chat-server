package chats

import (
	"context"

	"github.com/levon-dalakyan/chat-server/internal/converter"
	desc "github.com/levon-dalakyan/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.chatsService.SendMessage(ctx, converter.ToMessageCreateFromDesct(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
