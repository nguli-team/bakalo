package http

import (
	"bakalo.li/internal/config"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/repository"
	"bakalo.li/internal/service"
	"bakalo.li/internal/storage"
)

type ServiceContainer struct {
	BoardService  domain.BoardService
	ThreadService domain.ThreadService
}

func NewServiceContainer(cfg config.DatabaseConfig) (*ServiceContainer, error) {
	// storages
	gormDB, err := storage.NewGormPostgres(cfg)
	if err != nil {
		return nil, err
	}

	// services
	boardRepository := repository.NewGormBoardRepository(gormDB)
	threadRepository := repository.NewGormThreadRepository(gormDB)
	postRepository := repository.NewGormPostRepository(gormDB)

	// repositories
	boardService := service.NewBoardService(boardRepository)
	threadService := service.NewThreadService(threadRepository, postRepository)

	return &ServiceContainer{
		BoardService:  boardService,
		ThreadService: threadService,
	}, nil
}
