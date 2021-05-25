package service

import (
	"bakalo.li/internal/domain"
	"bakalo.li/internal/logger"
	"bakalo.li/internal/util"
	"context"
	"strconv"
)

type threadService struct {
	threadRepository domain.ThreadRepository
	postRepository   domain.PostRepository
}

func NewThreadService(tr domain.ThreadRepository, pr domain.PostRepository) domain.ThreadService {
	return &threadService{
		threadRepository: tr,
		postRepository:   pr,
	}
}

func (s threadService) FindByBoardID(ctx context.Context, boardID uint32) ([]*domain.Thread, error) {
	threads, err := s.threadRepository.FindByBoardID(ctx, boardID)
	if err != nil {
		return nil, err
	}
	return threads, err
}

func (s threadService) FindAll(ctx context.Context) ([]*domain.Thread, error) {
	threads, err := s.threadRepository.FindAll(ctx)
	for _, thread := range threads {
		op, err := s.postRepository.FindThreadOP(ctx, thread.ID)
		if err != nil {
			logger.Log.Warn(err)
			continue
		}
		thread.OPID = op.ID
		thread.OP = op
	}
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (s threadService) FindByID(ctx context.Context, id uint32) (*domain.Thread, error) {
	thread, err := s.threadRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (s threadService) Create(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	var err error
	thread, err = s.threadRepository.Create(ctx, thread)
	if err != nil {
		return nil, err
	}

	thread.OP.ThreadID = thread.ID

	tIDStr := strconv.FormatUint(uint64(thread.ID), 10)
	thread.OP.PosterID = util.GetMD5Hash(thread.OP.IPv4 + tIDStr)

	thread.OP, err = s.postRepository.Create(ctx, thread.OP)
	if err != nil {
		return nil, err
	}
	thread.OPID = thread.OP.ID

	return thread, nil
}

func (s threadService) Update(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	var err error
	thread, err = s.threadRepository.Update(ctx, thread)
	if err != nil {
		return nil, err
	}
	return thread, err
}
