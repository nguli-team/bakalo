package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/storage"
)

type gormPostRepository struct {
	DB *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) domain.PostRepository {
	return &gormPostRepository{
		DB: db,
	}
}

func (r gormPostRepository) FindAll(ctx context.Context) ([]domain.Post, error) {
	var posts []domain.Post
	result := r.DB.Find(&posts)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r gormPostRepository) FindByThreadID(
	ctx context.Context,
	threadID uint32,
) ([]domain.Post, error) {
	var posts []domain.Post
	result := r.DB.Where(&domain.Post{ThreadID: threadID}).Find(&posts)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r gormPostRepository) FindByID(ctx context.Context, id uint32) (*domain.Post, error) {
	var post *domain.Post
	result := r.DB.First(&post, id)
	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}
	return post, nil
}

func (r gormPostRepository) FindThreadOP(
	ctx context.Context,
	threadID uint32,
) (*domain.Post, error) {
	var post *domain.Post
	result := r.DB.Where(&domain.Post{ThreadID: threadID}).First(&post)
	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}
	return post, nil
}

func (r gormPostRepository) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	result := r.DB.Create(post)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r gormPostRepository) Update(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	panic("implement me")
}

func (r gormPostRepository) Delete(ctx context.Context, id uint32) error {
	r.DB.Delete(&domain.Post{}, id)
	return nil
}
