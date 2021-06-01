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
	TokenService  domain.TokenService
}

func NewServiceContainer(dbCfg config.DatabaseConfig, smtpCfg config.SMTPConfig) (*ServiceContainer, error) {
	// storages
	gormDB, err := persistence.NewGormPostgres(dbCfg)
	if err != nil {
		return nil, err
	}
	goCache := cache.NewGoCache()

	// repositories
	tokenRepository := repository.NewGormTokenRepository(gormDB)
	boardRepository := repository.NewGormBoardRepository(gormDB)
	threadRepository := repository.NewGormThreadRepository(gormDB)
	postRepository := repository.NewGormPostRepository(gormDB)

	// services
	tokenService := service.NewTokenService(tokenRepository, smtpCfg)
	boardService := service.NewBoardService(boardRepository, goCache)
	postService := service.NewPostService(postRepository, boardRepository, threadRepository, goCache)
	threadService := service.NewThreadService(threadRepository, postService, goCache)

	return &ServiceContainer{
		BoardService:  boardService,
		ThreadService: threadService,
		PostService:   postService,
		TokenService:  tokenService,
	}, nil
}
