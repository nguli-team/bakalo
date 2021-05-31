package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/storage"
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

func (r gormBoardRepository) FindByID(ctx context.Context, id uint32) (*domain.Board, error) {
	var board *domain.Board

	result := r.DB.First(&board, id)

	err := result.Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, storage.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return board, nil
}

func (r gormBoardRepository) FindByShorthand(
	ctx context.Context,
	shorthand string,
) (*domain.Board, error) {
	var board *domain.Board

	result := r.DB.Where(&domain.Board{Shorthand: shorthand}).First(&board)

	err := result.Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, storage.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return board, err
}

func (r gormBoardRepository) Create(
	ctx context.Context,
	board *domain.Board,
) (*domain.Board, error) {
	result := r.DB.Create(board)

	err := result.Error
	if err != nil {
		return nil, err
	}

	return board, nil
}

func (r gormBoardRepository) Update(
	ctx context.Context,
	board *domain.Board,
) (*domain.Board, error) {
	result := r.DB.Updates(board)

	err := result.Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, storage.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return board, err
}

func (r gormBoardRepository) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
