package domain

import (
	"context"
	"errors"
)

var (
	ErrTokenExpired = errors.New("token is expired")
	ErrTokenInvalid = errors.New("token is invalid")
)

type VIPToken struct {
	ID         uint32 `json:"id" gorm:"primaryKey"`
	IP         string `json:"ip" gorm:"index"`
	Token      string `json:"token" gorm:"uniqueIndex"`
	IsValid    bool   `json:"is_valid" gorm:"default:true"`
	ValidUntil uint32 `json:"valid_until"`
	CreatedAt  uint32 `json:"created_at"`
	UpdatedAt  uint32 `json:"updated_at"`
}

type TokenService interface {
	CheckVIPStatus(ctx context.Context, ip string) bool
	ValidateToken(ctx context.Context, token string) bool
	CreateNewToken(ctx context.Context, ip string) (*VIPToken, error)
	UpdateToken(ctx context.Context, token *VIPToken) (*VIPToken, error)
}

type TokenRepository interface {
	FindByIP(ctx context.Context, ip string) (*VIPToken, error)
	FindByToken(ctx context.Context, token string) (*VIPToken, error)
	Create(ctx context.Context, token *VIPToken) (*VIPToken, error)
	Update(ctx context.Context, token *VIPToken) (*VIPToken, error)
}
