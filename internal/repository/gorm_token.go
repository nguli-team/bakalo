package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/storage"
)

type gormTokenRepository struct {
	DB *gorm.DB
}

func NewGormTokenRepository(db *gorm.DB) domain.TokenRepository {
	return &gormTokenRepository{
		DB: db,
	}
}

func (r gormTokenRepository) FindByIP(ctx context.Context, ip string) (*domain.VIPToken, error) {
	var token *domain.VIPToken

	result := r.DB.Where(&domain.VIPToken{IP: ip}).First(&token)

	err := result.Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, storage.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return token, nil
}

func (r gormTokenRepository) FindByToken(ctx context.Context, token string) (*domain.VIPToken, error) {
	var tok *domain.VIPToken

	result := r.DB.Where(&domain.VIPToken{Token: token}).First(&tok)

	err := result.Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, storage.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return tok, nil
}

func (r gormTokenRepository) Create(ctx context.Context, token *domain.VIPToken) (*domain.VIPToken, error) {
	result := r.DB.Create(token)

	err := result.Error
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (r gormTokenRepository) Update(ctx context.Context, token *domain.VIPToken) (*domain.VIPToken, error) {
	result := r.DB.Updates(token)

	err := result.Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, storage.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return token, err
}
