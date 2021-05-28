package service

import (
	"context"
	"errors"
	"sort"
	"strconv"

	"bakalo.li/internal/domain"
	"bakalo.li/internal/util"
)

type postService struct {
	postRepository   domain.PostRepository
	boardRepository  domain.BoardRepository
	threadRepository domain.ThreadRepository
}

func NewPostService(
	postRepository domain.PostRepository,
	boardRepository domain.BoardRepository,
	threadRepository domain.ThreadRepository,
) domain.PostService {
	return &postService{
		postRepository:   postRepository,
		boardRepository:  boardRepository,
		threadRepository: threadRepository,
	}
}

func (s postService) FindAll(ctx context.Context) ([]domain.Post, error) {
	panic("implement me")
}

func (s postService) FindByThreadID(ctx context.Context, threadID uint32) ([]domain.Post, error) {
	posts, err := s.postRepository.FindByThreadID(ctx, threadID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s postService) FindThreadOP(ctx context.Context, threadID uint32) (*domain.Post, error) {
	op, err := s.postRepository.FindThreadOP(ctx, threadID)
	if err != nil {
		return nil, err
	}
	return op, nil
}

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

	// update board ref number
	board.RefCounter = newBoardRef
	_, err = s.boardRepository.Update(ctx, board)
	if err != nil {
		return nil, err
	}

	// fill post details
	post.PosterID = s.getPosterID(post.ThreadID, post.IPv4)
	post.Ref = newBoardRef

	// save the post
	post, err = s.postRepository.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	// update poster and media count
	threadPosts, err := s.FindByThreadID(ctx, post.ThreadID)
	if err != nil {
		return nil, err
	}
	postsCount := len(threadPosts)
	if postsCount == 0 {
		return nil, errors.New("failed to create post, no posts found in current thread")
	}
	var posters []string
	for _, threadPost := range threadPosts {
		posters = append(posters, threadPost.PosterID)
	}
	sort.Strings(posters)
	idx := sort.SearchStrings(posters, post.PosterID)
	// if poster not found in current thread (new poster)
	if idx != postsCount {
		thread.PosterCount = thread.PosterCount + 1
	}
	if post.MediaFileName != "" {
		thread.MediaCount = thread.MediaCount + 1
	}
	_, err = s.threadRepository.Update(ctx, thread)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s postService) getPosterID(threadID uint32, ip string) string {
	tIDStr := strconv.FormatUint(uint64(threadID), 10)
	posterID := util.GetMD5Hash(ip + tIDStr)[0:6]
	return posterID
}
