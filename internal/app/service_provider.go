package app

import (
	"context"
	"log"

	"github.com/levon-dalakyan/platform-common/pkg/closer"
	"github.com/levon-dalakyan/platform-common/pkg/db"
	"github.com/levon-dalakyan/platform-common/pkg/db/pg"

	chatsApi "github.com/levon-dalakyan/chat-server/internal/api/chats"
	"github.com/levon-dalakyan/chat-server/internal/config"
	"github.com/levon-dalakyan/chat-server/internal/repository"
	chatsRepository "github.com/levon-dalakyan/chat-server/internal/repository/chats"
	"github.com/levon-dalakyan/chat-server/internal/service"
	chatsService "github.com/levon-dalakyan/chat-server/internal/service/chats"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient        db.Client
	chatsRepository repository.ChatsRepository

	chatsService service.ChatsService

	chatsImpl *chatsApi.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) ChatsRepository(ctx context.Context) repository.ChatsRepository {
	if s.chatsRepository == nil {
		s.chatsRepository = chatsRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatsRepository
}

func (s *serviceProvider) ChatsService(ctx context.Context) service.ChatsService {
	if s.chatsService == nil {
		s.chatsService = chatsService.NewService(s.ChatsRepository(ctx))
	}

	return s.chatsService
}

func (s *serviceProvider) ChatsImpl(ctx context.Context) *chatsApi.Implementation {
	if s.chatsImpl == nil {
		s.chatsImpl = chatsApi.NewImplementation(s.ChatsService(ctx))
	}

	return s.chatsImpl
}
