package domain

import "context"

type Board struct {
	ID          uint32    `json:"id" gorm:"primaryKey"`
	Shorthand   string    `json:"shorthand" gorm:"uniqueIndex"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	RefCounter  uint32    `json:"ref_counter"`
	VipOnly     bool      `json:"vip_only"`
	CreatedAt   uint32    `json:"created_at"`
	UpdatedAt   uint32    `json:"updated_at"`
	Threads     []*Thread `json:"threads,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type BoardService interface {
	FindAll(ctx context.Context) ([]Board, error)
	FindByID(ctx context.Context, id uint32) (*Board, error)
	FindByShorthand(ctx context.Context, shorthand string) (*Board, error)
	Update(ctx context.Context, board *Board) (*Board, error)
}

type BoardRepository interface {
	FindAll(ctx context.Context) ([]Board, error)
	FindByID(ctx context.Context, id uint32) (*Board, error)
	FindByShorthand(ctx context.Context, shorthand string) (*Board, error)
	Create(ctx context.Context, board *Board) (*Board, error)
	Update(ctx context.Context, board *Board) (*Board, error)
	Delete(ctx context.Context, id int64) error
}
