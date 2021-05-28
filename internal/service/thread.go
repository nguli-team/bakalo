package service

import (
	"context"
	"errors"

	"bakalo.li/internal/domain"
	"bakalo.li/internal/logger"
	"bakalo.li/internal/storage"
)

type threadService struct {
	threadRepository domain.ThreadRepository
	postService      domain.PostService
}

func NewThreadService(
	threadRepository domain.ThreadRepository,
	postService domain.PostService,
) domain.ThreadService {
	return &threadService{
		threadRepository: threadRepository,
		postService:      postService,
	}
}

func (s threadService) fillOPDetails(ctx context.Context, thread *domain.Thread) error {
	op, err := s.postService.FindThreadOP(ctx, thread.ID)
	if err != nil {
		logger.Log().Warn(err)
		// if no OP is found, return without OP
		if errors.Is(err, storage.ErrRecordNotFound) {
			return nil
		}
		// else, pass the error
		return err
	}
	thread.OPID = op.ID
	thread.OP = op
	return nil
}

func (s threadService) FindByBoardID(ctx context.Context, boardID uint32) ([]domain.Thread, error) {
	options := &domain.ThreadsOptions{WithPosts: true}
	threads, err := s.threadRepository.FindByBoardID(ctx, boardID, options)
	if err != nil {
		return nil, err
	}
	return threads, err
}

func (s threadService) FindAll(ctx context.Context) ([]domain.Thread, error) {
	threads, err := s.threadRepository.FindAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	// TODO: use goroutine here
	for i, _ := range threads {
		err := s.fillOPDetails(ctx, &threads[i])
		if err != nil {
			continue
		}
	}

	return threads, nil
}

func (s threadService) FindByID(ctx context.Context, id uint32) (*domain.Thread, error) {
	thread, err := s.threadRepository.FindByID(ctx, id, nil)
	if err != nil {
		return nil, err
	}

	err = s.fillOPDetails(ctx, thread)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

func (s threadService) Create(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	thread, err := s.threadRepository.Create(ctx, thread)
	if err != nil {
		return nil, err
	}

	// fill in OP details
	op := thread.OP
	op.ThreadID = thread.ID

	// set poster and media count
	thread.PosterCount = 1
	thread.MediaCount = 1

	thread.OP, err = s.postService.Create(ctx, op)
	if err != nil {
		return nil, err
	}
	thread.OPID = thread.OP.ID

	return thread, nil
}

func (s threadService) Update(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	thread, err := s.threadRepository.Update(ctx, thread)
	if err != nil {
		return nil, err
	}

	// embed OP to thread before returning
	err = s.fillOPDetails(ctx, thread)
	if err != nil {
		return nil, err
	}

	return thread, err
}
