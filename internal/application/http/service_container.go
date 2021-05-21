package http

import (
	"bakalo.li/internal/config"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/repository"
	"bakalo.li/internal/service"
	"bakalo.li/internal/storage"
)

type ServiceContainer struct {
	BoardService domain.BoardService
}

func NewServiceContainer(cfg config.DatabaseConfig) (ServiceContainer, error) {
	// storages
	gormDB, err := storage.NewGormPostgres(cfg)
	if err != nil {
		return ServiceContainer{}, err
	}

	// board service
	boardRepository := repository.NewGormBoardRepository(gormDB)
	boardService := service.NewBoardService(boardRepository)

	return ServiceContainer{
		BoardService: boardService,
	}, nil
}
