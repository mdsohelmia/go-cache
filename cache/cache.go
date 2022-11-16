package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	strategy CacheStrategy
}

func NewCache(driver CacheStrategy) *Cache {
	return &Cache{
		strategy: driver,
	}
}
func (c *Cache) Get(key string) (any, error) {
	return nil, nil
}

func (c *Cache) Put(key string) (any, error) {
	return nil, nil
}
func (c *Cache) Has(key string) bool {
	return false
}
func (c Cache) Remember(key string, ttl time.Time, callback func()) any {
	fmt.Println("ok")
	return nil
}
func (c Cache) RememberForever(key string, callback CallbackFunc) any {
	return c.strategy.rememberForever(key, callback)
}
func (c Cache) Pull(key string) any {
	return nil
}
func (c Cache) Add(key string, value any, ttl ...time.Time) {}
func (c Cache) Forever(key string, value any)               {}
func (c Cache) Forget(key string) error {
	return nil
}
func (c Cache) Flush() string {
	return "flush"
}
