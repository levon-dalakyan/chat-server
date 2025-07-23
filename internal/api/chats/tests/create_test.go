package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/levon-dalakyan/chat-server/internal/api/chats"
	"github.com/levon-dalakyan/chat-server/internal/model"
	"github.com/levon-dalakyan/chat-server/internal/service"
	"github.com/levon-dalakyan/chat-server/internal/service/mocks"
	desc "github.com/levon-dalakyan/chat-server/pkg/chat_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatsServiceMockFunc func(mc *minimock.Controller) service.ChatsService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		usernames = []string{gofakeit.Username(), gofakeit.Username()}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Usernames: usernames,
		}

		chatCreateData = &model.ChatCreate{
			UserNames: usernames,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name             string
		args             args
		want             *desc.CreateResponse
		err              error
		chatsServiceMock chatsServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatsServiceMock: func(mc *minimock.Controller) service.ChatsService {
				mock := mocks.NewChatsServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatCreateData).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatsServiceMock: func(mc *minimock.Controller) service.ChatsService {
				mock := mocks.NewChatsServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatCreateData).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatsServiceMock := tt.chatsServiceMock(mc)
			api := chats.NewImplementation(chatsServiceMock)

			createRes, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, createRes)
		})
	}
}
