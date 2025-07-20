package chats

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/levon-dalakyan/chat-server/pkg/chat_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatsService.Delete(ctx, req.GetId())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
