package service

import (
	"context"

	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/storage/cache"
)

type boardService struct {
	boardRepository domain.BoardRepository
	cacheStorage    cache.Cache
}

func NewBoardService(
	boardRepository domain.BoardRepository,
	cacheStorage cache.Cache,
) domain.BoardService {
	return &boardService{
		boardRepository: boardRepository,
		cacheStorage:    cacheStorage,
	}
}

func (s boardService) FindAll(ctx context.Context) ([]domain.Board, error) {
	// check cache first
	cachedBoards, found := s.cacheStorage.Get(cache.AllBoardsKey)
	if found {
		return cachedBoards.([]domain.Board), nil
	}

	// get from DB
	boards, err := s.boardRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// cache DB Result
	s.cacheStorage.Set(cache.AllBoardsKey, boards, cache.DefaultExpiration)

	return boards, nil
}

func (s boardService) FindByID(ctx context.Context, id uint32) (*domain.Board, error) {
	board, err := s.boardRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (s boardService) FindByShorthand(
	ctx context.Context,
	shorthand string,
) (*domain.Board, error) {
	board, err := s.boardRepository.FindByShorthand(ctx, shorthand)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (s boardService) Update(ctx context.Context, board *domain.Board) (*domain.Board, error) {
	// invalidate caches first
	s.cacheStorage.Delete(cache.AllBoardsKey)

	board, err := s.boardRepository.Update(ctx, board)
	if err != nil {
		return nil, err
	}

	return board, nil
}
