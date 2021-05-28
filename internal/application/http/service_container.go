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
	PostService   domain.PostService
}

func NewServiceContainer(cfg config.DatabaseConfig) (*ServiceContainer, error) {
	// storages
	gormDB, err := storage.NewGormPostgres(cfg)

	if err != nil {
		return nil, err
	}

	// repositories
	boardRepository := repository.NewGormBoardRepository(gormDB)
	threadRepository := repository.NewGormThreadRepository(gormDB)
	postRepository := repository.NewGormPostRepository(gormDB)

	// services
	boardService := service.NewBoardService(boardRepository)
	postService := service.NewPostService(postRepository, boardRepository, threadRepository)
	threadService := service.NewThreadService(threadRepository, postService)

	return &ServiceContainer{
		BoardService:  boardService,
		ThreadService: threadService,
		PostService:   postService,
	}, nil
}
