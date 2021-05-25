package domain

import "context"

type Thread struct {
	ID          uint32  `json:"id" gorm:"primaryKey"`
	BoardID     uint32  `json:"board_id"`
	Title       string  `json:"title"`
	Sticky      bool    `json:"sticky"`
	Locked      bool    `json:"locked"`
	PosterCount uint32  `json:"poster_count"`
	MediaCount  uint32  `json:"media_count"`
	CreatedAt   uint32  `json:"created_at"`
	UpdatedAt   uint32  `json:"updated_at"`
	OP          *Post   `json:"op,omitempty" gorm:"-"`
	OPID        uint32  `json:"op_id,omitempty" gorm:"-"`
	Posts       []*Post `json:"posts,omitempty"`
}

type ThreadVM struct {
	*Thread
}

type ThreadService interface {
	FindAll(ctx context.Context) ([]*Thread, error)
	FindByBoardID(ctx context.Context, boardID uint32) ([]*Thread, error)
	FindByID(ctx context.Context, id uint32) (*Thread, error)
	Create(ctx context.Context, board *Thread) (*Thread, error)
	Update(ctx context.Context, board *Thread) (*Thread, error)
}

type ThreadRepository interface {
	FindAll(ctx context.Context) ([]*Thread, error)
	FindByBoardID(ctx context.Context, boardID uint32) ([]*Thread, error)
	FindByID(ctx context.Context, id uint32) (*Thread, error)
	Create(ctx context.Context, board *Thread) (*Thread, error)
	Update(ctx context.Context, board *Thread) (*Thread, error)
	Delete(ctx context.Context, id int64) error
}
