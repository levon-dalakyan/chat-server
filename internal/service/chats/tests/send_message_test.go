package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/levon-dalakyan/chat-server/internal/model"
	"github.com/levon-dalakyan/chat-server/internal/repository"
	"github.com/levon-dalakyan/chat-server/internal/repository/mocks"
	"github.com/levon-dalakyan/chat-server/internal/service/chats"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatsRepositoryMockFunc func(mc *minimock.Controller) repository.ChatsRepository

	type args struct {
		ctx     context.Context
		message *model.MessageCreate
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId    = gofakeit.Int64()
		from      = gofakeit.Name()
		text      = gofakeit.LoremIpsumSentence(20)
		timestamp = gofakeit.Date()

		repoErr = fmt.Errorf("repository error")

		message = &model.MessageCreate{
			ChatId:    chatId,
			From:      from,
			Text:      text,
			Timestamp: timestamp,
		}
	)

	tests := []struct {
		name                string
		args                args
		err                 error
		chatsRepositoryMock chatsRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:     ctx,
				message: message,
			},
			err: nil,
			chatsRepositoryMock: func(mc *minimock.Controller) repository.ChatsRepository {
				mock := mocks.NewChatsRepositoryMock(mc)
				mock.SendMessageMock.Expect(ctx, message).Return(nil)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx:     ctx,
				message: message,
			},
			err: repoErr,
			chatsRepositoryMock: func(mc *minimock.Controller) repository.ChatsRepository {
				mock := mocks.NewChatsRepositoryMock(mc)
				mock.SendMessageMock.Expect(ctx, message).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatsRepositoryMock := tt.chatsRepositoryMock(mc)
			service := chats.NewService(chatsRepositoryMock)

			err := service.SendMessage(tt.args.ctx, tt.args.message)
			require.Equal(t, tt.err, err)
		})
	}
}
