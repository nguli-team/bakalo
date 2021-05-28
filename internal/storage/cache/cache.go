package cache

import (
	"time"
)

const (
	AllBoardsKey  = "all-boards"
	AllThreadsKey = "all-threads"
	AllPostsKey   = "all-posts"

	BoardThreadsKeyPrefix = "twbid-"
	ThreadPostsKeyPrefix  = "pwtid-"

	// NoExpiration is for use with functions that take an expiration time.
	NoExpiration time.Duration = -1

	// DefaultExpiration is for use with functions that take an expiration time.
	DefaultExpiration time.Duration = 0
)

type Cache interface {
	Set(key string, data interface{}, expiration time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
}
