package http

import (
	"github.com/nguli-team/bakalo/internal/config"
	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/repository"
	"github.com/nguli-team/bakalo/internal/service"
	"github.com/nguli-team/bakalo/internal/storage/cache"
	"github.com/nguli-team/bakalo/internal/storage/persistence"
)

type ServiceContainer struct {
	BoardService  domain.BoardService
	ThreadService domain.ThreadService
	PostService   domain.PostService
}

func NewServiceContainer(cfg config.DatabaseConfig) (*ServiceContainer, error) {
	// storages
	gormDB, err := persistence.NewGormPostgres(cfg)
	if err != nil {
		return nil, err
	}
	goCache := cache.NewGoCache()

	// repositories
	boardRepository := repository.NewGormBoardRepository(gormDB)
	threadRepository := repository.NewGormThreadRepository(gormDB)
	postRepository := repository.NewGormPostRepository(gormDB)

	// services
	boardService := service.NewBoardService(boardRepository, goCache)
	postService := service.NewPostService(postRepository, boardRepository, threadRepository, goCache)
	threadService := service.NewThreadService(threadRepository, postService, goCache)

	return &ServiceContainer{
		BoardService:  boardService,
		ThreadService: threadService,
		PostService:   postService,
	}, nil
}
