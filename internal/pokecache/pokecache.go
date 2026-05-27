package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entry: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}
	// newCache.Add(key, value)

	// interval := time.Duration(2) * time.Second
	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	elem, ok := c.entry[key]
	return elem.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	// defer ticker.Stop()
	for range ticker.C {
		c.reap(c.entry, interval)
	}
}

func (c *Cache) reap(t map[string]cacheEntry, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range t {
		if time.Since(val.createdAt) > interval {
			delete(c.entry, key)
		}
	}
}
