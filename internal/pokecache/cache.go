package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache      map[string]cacheEntry
	mu         *sync.Mutex
	expireTime time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *cacheEntry) isExpired(expireTime time.Duration) bool {
	return c.createdAt.Before(time.Now().Add(-expireTime))
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}

	if entry.isExpired(c.expireTime) {
		delete(c.cache, key)
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.cache {
		if entry.isExpired(last) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) realLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func NewCache(expireTime, interval time.Duration) *Cache {
	cache := &Cache{
		cache:      make(map[string]cacheEntry),
		mu:         &sync.Mutex{},
		expireTime: expireTime,
	}

	go cache.realLoop(interval)

	return cache
}
