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

func (r gormBoardRepository) FindAll(ctx context.Context) ([]domain.Board, error) {
	var boards []domain.Board
	result := r.DB.Find(&boards)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (r gormBoardRepository) FindByID(ctx context.Context, id int64) (domain.Board, error) {
	panic("implement me")
}

func (r gormBoardRepository) FindByShorthand(ctx context.Context, shorthand string) (domain.Board, error) {
	panic("implement me")
}

func (r gormBoardRepository) Create(ctx context.Context, board *domain.Board) error {
	result := r.DB.Create(&board)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (r gormBoardRepository) Update(ctx context.Context, board *domain.Board) error {
	panic("implement me")
}

func (r gormBoardRepository) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
