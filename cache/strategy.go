package cache

import (
	"time"
)

type CallbackFunc func() any

type CacheStrategy interface {
	get(key string) (any, error)
	has(key string) bool
	remember(key string, ttl time.Time, callback CallbackFunc) any
	rememberForever(key string, callback CallbackFunc) any
	pull(key string) any
	put(key string, value any, ttl ...time.Time)
	add(key string, value any, ttl ...time.Time)
	forever(key string, value any)
	forget(key string) error
	flush() string
}
