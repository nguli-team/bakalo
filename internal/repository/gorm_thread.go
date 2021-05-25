package repository

import (
	"bakalo.li/internal/domain"
	"bakalo.li/internal/storage"
	"context"
	"errors"
	"gorm.io/gorm"
)

type gormThreadRepository struct {
	DB *gorm.DB
}

func NewGormThreadRepository(db *gorm.DB) domain.ThreadRepository {
	return &gormThreadRepository{
		DB: db,
	}
}

func (r gormThreadRepository) FindAll(ctx context.Context) ([]*domain.Thread, error) {
	var threads []*domain.Thread
	result := r.DB.Find(&threads)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (r gormThreadRepository) FindByBoardID(ctx context.Context, boardID uint32) ([]*domain.Thread, error) {
	var threads []*domain.Thread
	result := r.DB.Where(&domain.Thread{BoardID: boardID}).Find(threads)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (r gormThreadRepository) FindByID(ctx context.Context, id uint32) (*domain.Thread, error) {
	var thread *domain.Thread
	result := r.DB.Find(thread, id)
	err := result.Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}
	return thread, nil
}

func (r gormThreadRepository) Create(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	result := r.DB.Create(thread)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (r gormThreadRepository) Update(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	result := r.DB.Save(thread)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (r gormThreadRepository) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
