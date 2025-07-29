package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ldevprog/chat-server/internal/api/chats"
	"github.com/ldevprog/chat-server/internal/service"
	"github.com/ldevprog/chat-server/internal/service/mocks"
	desc "github.com/ldevprog/chat-server/pkg/chat_v1"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatsServiceMockFunc func(mc *minimock.Controller) service.ChatsService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &desc.DeleteRequest{
			Id: id,
		}
	)

	tests := []struct {
		name             string
		args             args
		want             *emptypb.Empty
		err              error
		chatsServiceMock chatsServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			chatsServiceMock: func(mc *minimock.Controller) service.ChatsService {
				mock := mocks.NewChatsServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  serviceErr,
			chatsServiceMock: func(mc *minimock.Controller) service.ChatsService {
				mock := mocks.NewChatsServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatsServiceMock := tt.chatsServiceMock(mc)
			api := chats.NewImplementation(chatsServiceMock)

			res, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
