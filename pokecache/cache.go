package pokecache

import "time"

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mutex    chan struct{}
	NewCache func(time.Duration) *Cache
}

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		entries: make(map[string]cacheEntry),
		mutex:   make(chan struct{}, 1),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex <- struct{}{}
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	<-c.mutex
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex <- struct{}{}
	entry, ok := c.entries[key]
	<-c.mutex
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mutex <- struct{}{}
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		<-c.mutex
	}
}
