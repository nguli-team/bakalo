package repository

import (
	"bakalo.li/internal/domain"
	"bakalo.li/internal/storage"
	"context"
	"errors"
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

	return boards, err
}

func (r gormBoardRepository) FindByID(ctx context.Context, id int64) (domain.Board, error) {
	var board domain.Board

	result := r.DB.First(&board, id)
	err := result.Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return board, storage.ErrRecordNotFound
	}

	if err != nil {
		return board, err
	}

	return board, nil
}

func (r gormBoardRepository) FindByShorthand(ctx context.Context, shorthand string) (domain.Board, error) {
	var board domain.Board

	result := r.DB.Where(&domain.Board{Shorthand: shorthand}).First(&board)
	err := result.Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return board, storage.ErrRecordNotFound
	}

	if err != nil {
		return board, nil
	}

	return board, err
}

func (r gormBoardRepository) Create(ctx context.Context, board *domain.Board) (domain.Board, error) {
	// FIXME: this is a horrible implementation
	result := r.DB.Create(board)
	err := result.Error
	if err != nil {
		return domain.Board{}, err
	}

	return *board, nil
}

func (r gormBoardRepository) Update(ctx context.Context, board *domain.Board) (domain.Board, error) {
	panic("implement me")
}

func (r gormBoardRepository) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
