package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ldevprog/chat-server/internal/api/chats"
	"github.com/ldevprog/chat-server/internal/model"
	"github.com/ldevprog/chat-server/internal/service"
	"github.com/ldevprog/chat-server/internal/service/mocks"
	desc "github.com/ldevprog/chat-server/pkg/chat_v1"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatsServiceMockFunc func(mc *minimock.Controller) service.ChatsService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId    = gofakeit.Int64()
		from      = gofakeit.Name()
		text      = gofakeit.LoremIpsumParagraph(1, 2, 20, ".")
		timestamp = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &desc.SendMessageRequest{
			ChatId:    chatId,
			From:      from,
			Text:      text,
			Timestamp: timestamppb.New(timestamp),
		}

		message = &model.MessageCreate{
			ChatId:    chatId,
			From:      from,
			Text:      text,
			Timestamp: timestamp,
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
				mock.SendMessageMock.Expect(ctx, message).Return(nil)
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
				mock.SendMessageMock.Expect(ctx, message).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatsServiceMock := tt.chatsServiceMock(mc)
			api := chats.NewImplementation(chatsServiceMock)

			res, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
