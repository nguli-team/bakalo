package service

import (
	"context"
	"errors"

	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/logger"
	"github.com/nguli-team/bakalo/internal/storage"
	"github.com/nguli-team/bakalo/internal/storage/cache"
	"github.com/nguli-team/bakalo/internal/util"
)

type threadService struct {
	threadRepository domain.ThreadRepository
	postService      domain.PostService
	cacheStorage     cache.Cache
}

func NewThreadService(
	threadRepository domain.ThreadRepository,
	postService domain.PostService,
	cacheStorage cache.Cache,
) domain.ThreadService {
	return &threadService{
		threadRepository: threadRepository,
		postService:      postService,
		cacheStorage:     cacheStorage,
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
	cacheKey := cache.BoardThreadsKeyPrefix + util.Uint32ToStr(boardID)

	// check cache first
	cachedThreads, found := s.cacheStorage.Get(cacheKey)
	if found {
		return cachedThreads.([]domain.Thread), nil
	}

	// get from DB
	options := &domain.ThreadsOptions{WithPosts: false}
	threads, err := s.threadRepository.FindByBoardID(ctx, boardID, options)
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

	// cache DB result
	s.cacheStorage.Set(cacheKey, threads, cache.DefaultExpiration)

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
	cacheKey := cache.ThreadPostsKeyPrefix + util.Uint32ToStr(id)
	options := &domain.ThreadsOptions{WithPosts: true}

	cachedPosts, found := s.cacheStorage.Get(cacheKey)
	if found {
		options.WithPosts = false
	}

	thread, err := s.threadRepository.FindByID(ctx, id, options)
	if err != nil {
		return nil, err
	}

	if found {
		thread.Posts = cachedPosts.([]domain.Post)
	} else {
		s.cacheStorage.Set(cacheKey, thread.Posts, cache.DefaultExpiration)
	}

	err = s.fillOPDetails(ctx, thread)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

func (s threadService) Create(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	// invalidate caches first
	s.cacheStorage.Delete(cache.AllThreadsKey)
	s.cacheStorage.Delete(cache.BoardThreadsKeyPrefix + util.Uint32ToStr(thread.BoardID))

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
	// invalidate caches first
	s.cacheStorage.Delete(cache.AllThreadsKey)
	s.cacheStorage.Delete(cache.BoardThreadsKeyPrefix + util.Uint32ToStr(thread.BoardID))

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
