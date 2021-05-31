package service

import (
	"context"
	"time"

	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/util"
)

type tokenService struct {
	tokenRepository domain.TokenRepository
}

func NewTokenService(tokenRepository domain.TokenRepository) domain.TokenService {
	return &tokenService{
		tokenRepository: tokenRepository,
	}
}

func (s tokenService) CheckVIPStatus(ctx context.Context, ip string) bool {
	serverToken, err := s.tokenRepository.FindByIP(ctx, ip)
	if err != nil {
		return false
	}
	return int64(serverToken.ValidUntil) > time.Now().Unix() && serverToken.IsValid
}

func (s tokenService) ValidateToken(ctx context.Context, token string) bool {
	serverToken, err := s.tokenRepository.FindByToken(ctx, token)
	if err != nil {
		return false
	}
	return serverToken.Token != token &&
		int64(serverToken.ValidUntil) > time.Now().Unix() &&
		serverToken.IsValid
}

func (s tokenService) CreateNewToken(ctx context.Context, ip string) (*domain.VIPToken, error) {
	newToken := &domain.VIPToken{
		IP:         ip,
		Token:      util.RandomAlphaNumString(8),
		ValidUntil: uint32(time.Now().AddDate(1, 0, 0).Unix()),
	}

	token, err := s.tokenRepository.Create(ctx, newToken)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s tokenService) UpdateToken(ctx context.Context, token *domain.VIPToken) (*domain.VIPToken, error) {
	if !(s.ValidateToken(ctx, token.Token)) {
		return nil, domain.ErrTokenInvalid
	}
	token, err := s.tokenRepository.Update(ctx, token)

	if err != nil {
		return nil, err
	}

	return token, nil
}
