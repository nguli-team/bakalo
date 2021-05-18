package domain

import (
	"context"
	"time"
)

type Board struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Shorthand   string    `json:"shorthand" gorm:"uniqueIndex"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	RefCounter  int64     `json:"ref_counter"`
	VipOnly     bool      `json:"vip_only"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BoardService interface {
}

type BoardRepository interface {
	FindAll(ctx context.Context) ([]Board, error)
	FindByID(ctx context.Context, id int64) (Board, error)
	FindByShorthand(ctx context.Context, shorthand string) (Board, error)
	Create(ctx context.Context, board *Board) error
	Update(ctx context.Context, board *Board) error
	Delete(ctx context.Context, id int64) error
}
