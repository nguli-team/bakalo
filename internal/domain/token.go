package domain

import (
	"context"
	"errors"
)

var ErrTokenExpired = errors.New("token is expired")
var ErrTokenInvalid = errors.New("token is invalid")

type VIPToken struct {
	ID         uint32 `json:"id" gorm:"primaryKey"`
	IP         string `json:"ip"`
	Token      string `json:"token"`
	IsValid    bool   `json:"is_valid"`
	ValidUntil uint32 `json:"valid_until"`
	CreatedAt  uint32 `json:"created_at"`
	UpdatedAt  uint32 `json:"updated_at"`
}

type TokenService interface {
	ValidateToken(ctx context.Context, token *VIPToken) bool
	CreateNewToken(ctx context.Context, token *VIPToken) (*VIPToken, error)
	UpdateToken(ctx context.Context, token *VIPToken) (*VIPToken, error)
}

type TokenRepository interface {
	FindByIP(ctx context.Context, ip string) (*VIPToken, error)
	Create(ctx context.Context, token *VIPToken) (*VIPToken, error)
	Update(ctx context.Context, token *VIPToken) (*VIPToken, error)
}
