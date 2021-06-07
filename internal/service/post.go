package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/storage/cache"
	"github.com/nguli-team/bakalo/internal/util"
)

const (
	createPost = iota
	deletePost
)

type postService struct {
	postRepository   domain.PostRepository
	boardRepository  domain.BoardRepository
	threadRepository domain.ThreadRepository
	cacheStorage     cache.Cache
}

func NewPostService(
	postRepository domain.PostRepository,
	boardRepository domain.BoardRepository,
	threadRepository domain.ThreadRepository,
	cacheStorage cache.Cache,
) domain.PostService {
	return &postService{
		postRepository:   postRepository,
		boardRepository:  boardRepository,
		threadRepository: threadRepository,
		cacheStorage:     cacheStorage,
	}
}

// FindAll ...
func (s postService) FindAll(ctx context.Context) ([]domain.Post, error) {
	// check cache first
	cachedPosts, found := s.cacheStorage.Get(cache.AllPostsKey)
	if found {
		return cachedPosts.([]domain.Post), nil
	}

	// get from DB
	posts, err := s.postRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// cache DB result
	s.cacheStorage.Set(cache.AllPostsKey, posts, cache.DefaultExpiration)

	return posts, nil
}

func (s postService) FindByID(ctx context.Context, id uint32) (*domain.Post, error) {
	post, err := s.postRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// FindByThreadID ...
func (s postService) FindByThreadID(ctx context.Context, threadID uint32) ([]domain.Post, error) {
	posts, err := s.postRepository.FindByThreadID(ctx, threadID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// FindThreadOP ...
func (s postService) FindThreadOP(ctx context.Context, threadID uint32) (*domain.Post, error) {
	op, err := s.postRepository.FindThreadOP(ctx, threadID)
	if err != nil {
		return nil, err
	}
	return op, nil
}

// Create ...
func (s postService) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	// validation
	if post.IPv4 == "" {
		return nil, errors.New("poster IP is empty")
	}

	// get ref number
	thread, err := s.threadRepository.FindByID(
		ctx,
		post.ThreadID,
		&domain.ThreadsOptions{WithPosts: false},
	)
	if err != nil {
		return nil, err
	}
	board, err := s.boardRepository.FindByID(ctx, thread.BoardID)
	if err != nil {
		return nil, err
	}
	newBoardRef := board.RefCounter + 1

	// invalidate caches
	s.cacheStorage.Delete(cache.AllPostsKey)
	s.cacheStorage.Delete(cache.ThreadPostsKeyPrefix + util.Uint32ToStr(thread.ID))

	// update board ref number
	board.RefCounter = newBoardRef
	_, err = s.boardRepository.Update(ctx, board)
	if err != nil {
		return nil, err
	}

	// fill post details
	post.PosterID = s.getPosterID(post.ThreadID, post.IPv4)
	post.Ref = newBoardRef

	// update poster and media count
	err = s.updateThreadInfo(ctx, post, thread, createPost)
	if err != nil {
		return nil, err
	}

	// save the post
	post, err = s.postRepository.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s postService) Delete(ctx context.Context, id uint32) error {
	post, err := s.postRepository.FindByID(ctx, id)
	if err != nil {
		return nil
	}

	threadOP, err := s.FindThreadOP(ctx, post.ThreadID)
	if err == nil {
		if threadOP.ID == id {
			return errors.New("post is an OP and cannot be deleted, delete the thread instead")
		}
	}

	thread, err := s.threadRepository.FindByID(ctx, post.ThreadID, nil)
	if err != nil {
		return err
	}

	err = s.updateThreadInfo(ctx, post, thread, deletePost)
	if err != nil {
		return err
	}

	// invalidate caches
	s.cacheStorage.Delete(cache.AllPostsKey)
	s.cacheStorage.Delete(cache.ThreadPostsKeyPrefix + util.Uint32ToStr(thread.ID))

	_ = s.postRepository.Delete(ctx, id)

	return nil
}

// updateThreadInfo ...
func (s postService) updateThreadInfo(
	ctx context.Context,
	post *domain.Post,
	thread *domain.Thread,
	op int,
) error {
	threadPosts, err := s.FindByThreadID(ctx, post.ThreadID)
	if err != nil {
		return err
	}

	// get posters in thread
	var posters []string
	for _, threadPost := range threadPosts {
		posters = append(posters, threadPost.PosterID)
	}

	// find current poster in thread posters
	newPoster := !util.ContainsString(posters, post.PosterID)

	switch op {
	case createPost:
		// if new poster
		if newPoster {
			thread.PosterCount = thread.PosterCount + 1
		}
		// update media count
		if post.MediaFileName != "" {
			thread.MediaCount = thread.MediaCount + 1
		}
	case deletePost:
		// if new poster
		if newPoster {
			thread.PosterCount = thread.PosterCount - 1
		}
		// update media count
		if post.MediaFileName != "" {
			thread.MediaCount = thread.MediaCount - 1
		}
	}

	// invalidate caches
	s.cacheStorage.Delete(cache.AllThreadsKey)
	s.cacheStorage.Delete(cache.BoardThreadsKeyPrefix + util.Uint32ToStr(thread.BoardID))

	// update thread
	_, err = s.threadRepository.Update(ctx, thread)
	if err != nil {
		return err
	}

	return nil
}

// getPosterID ...
func (s postService) getPosterID(threadID uint32, ip string) string {
	tIDStr := strconv.FormatUint(uint64(threadID), 10)
	posterID := util.GetMD5Hash(ip + tIDStr)[0:6]
	return posterID
}
