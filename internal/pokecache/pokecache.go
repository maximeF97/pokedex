package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		interval: interval,
		entries:  make(map[string]cacheEntry),
	}
	go c.reapLoop()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	s := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = s
	c.mu.Unlock()
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	value, exists := c.entries[key]
	c.mu.Unlock()

	if exists {
		return value.val, true
	} else {
		return nil, false
	}
}
func (c *Cache) reapLoop() {

	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			t := time.Since(entry.createdAt)
			if t > c.interval {
				delete(c.entries, key)

			}

		}
		c.mu.Unlock()
	}

}
