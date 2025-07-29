package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ldevprog/chat-server/internal/repository"
	"github.com/ldevprog/chat-server/internal/repository/mocks"
	"github.com/ldevprog/chat-server/internal/service/chats"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatsRepositoryMockFunc func(mc *minimock.Controller) repository.ChatsRepository

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		repoErr = fmt.Errorf("repository error")
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
				ctx: ctx,
				id:  id,
			},
			err: nil,
			chatsRepositoryMock: func(mc *minimock.Controller) repository.ChatsRepository {
				mock := mocks.NewChatsRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			err: repoErr,
			chatsRepositoryMock: func(mc *minimock.Controller) repository.ChatsRepository {
				mock := mocks.NewChatsRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatsRepositoryMock := tt.chatsRepositoryMock(mc)
			service := chats.NewService(chatsRepositoryMock)

			err := service.Delete(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
		})
	}
}
