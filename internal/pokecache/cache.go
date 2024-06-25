package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.Mutex
	duration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	// @TODO implement Cache Add
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// @TODO implement Cache Get
	return []byte{}, false
}

func (c *Cache) realLoop() {
	// @TODO implement cache real loop
}

func NewCache() *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry),
	}

	cache.realLoop()

	return cache
}
