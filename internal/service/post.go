package service

import (
	"bakalo.li/internal/domain"
	"context"
)

type postService struct {
	repository domain.PostRepository
}

func NewPostService(repository domain.PostRepository) domain.PostService {
	return &postService{
		repository: repository,
	}
}

func (s postService) FindAll(ctx context.Context) ([]*domain.Post, error) {
	panic("implement me")
}

func (s postService) FindThreadOP(ctx context.Context, threadID uint32) (*domain.Post, error) {
	op, err := s.repository.FindThreadOP(ctx, threadID)
	if err != nil {
		return nil, err
	}
	return op, nil
}
