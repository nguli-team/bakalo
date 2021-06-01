package domain

import "context"

type Post struct {
	ID            uint32 `json:"id" gorm:"primaryKey"`
	Ref           uint32 `json:"ref" gorm:"index"`
	ThreadID      uint32 `json:"thread_id"`
	PosterID      string `json:"poster_id"`
	IPv4          string `json:"-"`
	Name          string `json:"name,omitempty"`
	Text          string `json:"text"`
	MediaFileName string `json:"media_file_name,omitempty"`
	CreatedAt     uint32 `json:"created_at"`
	UpdatedAt     uint32 `json:"updated_at"`
}

type PostService interface {
	FindAll(ctx context.Context) ([]Post, error)
	FindByThreadID(ctx context.Context, threadID uint32) ([]Post, error)
	FindThreadOP(ctx context.Context, threadID uint32) (*Post, error)
	FindByID(ctx context.Context, id uint32) (*Post, error)
	Create(ctx context.Context, post *Post) (*Post, error)
	Delete(ctx context.Context, id uint32) error
}

type PostRepository interface {
	FindAll(ctx context.Context) ([]Post, error)
	FindByThreadID(ctx context.Context, threadID uint32) ([]Post, error)
	FindByID(ctx context.Context, id uint32) (*Post, error)
	FindThreadOP(ctx context.Context, threadID uint32) (*Post, error)
	Create(ctx context.Context, post *Post) (*Post, error)
	Update(ctx context.Context, post *Post) (*Post, error)
	Delete(ctx context.Context, id uint32) error
}
