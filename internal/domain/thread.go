package domain

import "context"

type Thread struct {
	ID          uint32 `json:"id" gorm:"primaryKey"`
	BoardID     uint32 `json:"board_id"`
	Title       string `json:"title"`
	Sticky      bool   `json:"sticky"`
	Locked      bool   `json:"locked"`
	PosterCount uint32 `json:"poster_count"`
	ReplyCount  uint32 `json:"reply_count" gorm:"-"`
	MediaCount  uint32 `json:"media_count"`
	CreatedAt   uint32 `json:"created_at"`
	UpdatedAt   uint32 `json:"updated_at"`
	OPID        uint32 `json:"op_id,omitempty" gorm:"-"`
	OP          *Post  `json:"op,omitempty" gorm:"-"`
	Posts       []Post `json:"posts,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ThreadsOptions struct {
	WithPosts bool
	Limit     int
	Offset    int
}

type ThreadService interface {
	FindAll(ctx context.Context) ([]Thread, error)
	FindPopular(ctx context.Context) ([]Thread, error)
	FindByBoardID(ctx context.Context, boardID uint32) ([]Thread, error)
	FindByID(ctx context.Context, id uint32) (*Thread, error)
	Create(ctx context.Context, board *Thread) (*Thread, error)
	Update(ctx context.Context, board *Thread) (*Thread, error)
	Delete(ctx context.Context, id uint32) error
}

type ThreadRepository interface {
	FindAll(ctx context.Context, options *ThreadsOptions) ([]Thread, error)
	FindPopular(ctx context.Context, options *ThreadsOptions) ([]Thread, error)
	FindByBoardID(ctx context.Context, boardID uint32, options *ThreadsOptions) ([]Thread, error)
	FindByID(ctx context.Context, id uint32, options *ThreadsOptions) (*Thread, error)
	Create(ctx context.Context, board *Thread) (*Thread, error)
	Update(ctx context.Context, board *Thread) (*Thread, error)
	Delete(ctx context.Context, id uint32) error
}
