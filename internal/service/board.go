package service

import (
	"context"

	"bakalo.li/internal/domain"
)

type boardService struct {
	repository domain.BoardRepository
}

func NewBoardService(repository domain.BoardRepository) domain.BoardService {
	return &boardService{
		repository: repository,
	}
}

func (s boardService) FindAll(ctx context.Context) ([]domain.Board, error) {
	boards, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return boards, nil
}

func (s boardService) FindByID(ctx context.Context, id uint32) (*domain.Board, error) {
	board, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (s boardService) FindByShorthand(
	ctx context.Context,
	shorthand string,
) (*domain.Board, error) {
	board, err := s.repository.FindByShorthand(ctx, shorthand)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (s boardService) Update(ctx context.Context, board *domain.Board) (*domain.Board, error) {
	board, err := s.repository.Update(ctx, board)
	if err != nil {
		return nil, err
	}
	return board, nil
}
