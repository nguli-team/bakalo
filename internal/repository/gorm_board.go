package repository

import (
	"bakalo.li/internal/domain"
	"context"
	"gorm.io/gorm"
)

type gormBoardRepository struct {
	DB *gorm.DB
}

func NewGormBoardRepository(db *gorm.DB) domain.BoardRepository {
	return &gormBoardRepository{
		DB: db,
	}
}

func (repository gormBoardRepository) FindAll(ctx context.Context) ([]domain.Board, error) {
	var boards []domain.Board
	result := repository.DB.Find(&boards)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (repository gormBoardRepository) FindByID(ctx context.Context, id int64) (domain.Board, error) {
	panic("implement me")
}

func (repository gormBoardRepository) FindByShorthand(ctx context.Context, shorthand string) (domain.Board, error) {
	panic("implement me")
}

func (repository gormBoardRepository) Create(ctx context.Context, board *domain.Board) error {
	result := repository.DB.Create(&board)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (repository gormBoardRepository) Update(ctx context.Context, board *domain.Board) error {
	panic("implement me")
}

func (repository gormBoardRepository) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
