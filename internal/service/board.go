package service

import (
	"bakalo.li/internal/domain"
	"context"
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

func (s boardService) FindByID(ctx context.Context, id int64) (domain.Board, error) {
	panic("implement me")
}

func (s boardService) FindByShorthand(ctx context.Context, shorthand string) (domain.Board, error) {
	panic("implement me")
}
