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

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatsRepositoryMockFunc func(mc *minimock.Controller) repository.ChatsRepository

	type args struct {
		ctx            context.Context
		chatCreateData *model.ChatCreate
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		usernames = []string{gofakeit.Username(), gofakeit.Username()}

		repoErr = fmt.Errorf("repository error")

		chatCreateData = &model.ChatCreate{
			UserNames: usernames,
		}
	)

	tests := []struct {
		name                string
		args                args
		want                int64
		err                 error
		chatsRepositoryMock chatsRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:            ctx,
				chatCreateData: chatCreateData,
			},
			want: id,
			err:  nil,
			chatsRepositoryMock: func(mc *minimock.Controller) repository.ChatsRepository {
				mock := mocks.NewChatsRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, chatCreateData).Return(id, nil)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx:            ctx,
				chatCreateData: chatCreateData,
			},
			want: 0,
			err:  repoErr,
			chatsRepositoryMock: func(mc *minimock.Controller) repository.ChatsRepository {
				mock := mocks.NewChatsRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, chatCreateData).Return(0, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatsRepositoryMock := tt.chatsRepositoryMock(mc)
			service := chats.NewService(chatsRepositoryMock)

			res, err := service.Create(tt.args.ctx, tt.args.chatCreateData)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
