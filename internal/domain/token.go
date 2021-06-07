package domain

import (
	"context"
	"errors"
	"net/http"
)

var (
	ErrTokenExpired = errors.New("token is expired")
	ErrTokenInvalid = errors.New("token is invalid")
	AdminContextKey = "admin"
	VipContextKey   = "vip"
)

type VipToken struct {
	ID         uint32 `json:"id" gorm:"primaryKey"`
	IP         string `json:"ip" gorm:"index"`
	Token      string `json:"token" gorm:"uniqueIndex"`
	PIN        int    `json:"-"`
	Email      string `json:"email"`
	IsValid    bool   `json:"is_valid" gorm:"default:true"`
	IsAdmin    bool   `json:"is_admin" gorm:"default:false"`
	ValidUntil uint32 `json:"valid_until"`
	CreatedAt  uint32 `json:"created_at"`
	UpdatedAt  uint32 `json:"updated_at"`
}

type TokenService interface {
	FindByIP(ctx context.Context, ip string) (*VipToken, error)
	CheckVIPStatus(ctx context.Context, ip string) bool
	ValidateToken(ctx context.Context, token string, pin int) bool
	CreateNewToken(ctx context.Context, ip string, email string) (*VipToken, error)
	UpdateTokenIP(ctx context.Context, ip string, token string) (*VipToken, error)
	SendTokenViaEmail(ctx context.Context, token *VipToken) error
}

type TokenRepository interface {
	FindByIP(ctx context.Context, ip string) (*VipToken, error)
	FindByToken(ctx context.Context, token string) (*VipToken, error)
	Create(ctx context.Context, token *VipToken) (*VipToken, error)
	Update(ctx context.Context, token *VipToken) (*VipToken, error)
}

func IsAdminRequest(r *http.Request) bool {
	ctx := r.Context()

	if isAdmin, ok := ctx.Value(AdminContextKey).(bool); ok {
		return isAdmin
	}

	return false
}
