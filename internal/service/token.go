package service

import (
	"context"

	"github.com/nguli-team/bakalo/internal/domain"
)

type tokenService struct {
	tokenRepository domain.TokenRepository
}

func NewTokenService(tokenRepository domain.TokenRepository) domain.TokenService {
	return &tokenService{
		tokenRepository: tokenRepository,
	}
}

func (t tokenService) ValidateToken(ctx context.Context, token *domain.VIPToken) bool {
	panic("implement me")
}

func (t tokenService) CreateNewToken(ctx context.Context, token *domain.VIPToken) (*domain.VIPToken, error) {
	panic("implement me")
}

func (t tokenService) UpdateToken(ctx context.Context, token *domain.VIPToken) (*domain.VIPToken, error) {
	panic("implement me")
}
