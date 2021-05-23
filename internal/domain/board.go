package domain

import (
	"context"
)

type Board struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Shorthand   string `json:"shorthand" gorm:"uniqueIndex"`
	Name        string `json:"name"`
	Description string `json:"description"`
	RefCounter  int64  `json:"ref_counter"`
	VipOnly     bool   `json:"vip_only"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type BoardService interface {
	FindAll(ctx context.Context) ([]Board, error)
	FindByID(ctx context.Context, id int64) (Board, error)
	FindByShorthand(ctx context.Context, shorthand string) (Board, error)
}

type BoardRepository interface {
	FindAll(ctx context.Context) ([]Board, error)
	FindByID(ctx context.Context, id int64) (Board, error)
	FindByShorthand(ctx context.Context, shorthand string) (Board, error)
	Create(ctx context.Context, board *Board) (Board, error)
	Update(ctx context.Context, board *Board) (Board, error)
	Delete(ctx context.Context, id int64) error
}
