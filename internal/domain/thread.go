package domain

type Thread struct {
	ID        int64 `json:"id" gorm:"primaryKey"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
