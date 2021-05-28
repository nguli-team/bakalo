package cache

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
)

func NewGoCache() Cache {
	c := goCache.New(5*time.Minute, 10*time.Minute)
	return c
}
