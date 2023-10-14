package cache

import (
	"sync"
	"time"
)

type Cache struct {
	store sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.store.Store(key, value)
	time.AfterFunc(ttl, func() {
		c.store.Delete(key)
	})
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.store.Load(key)
}
