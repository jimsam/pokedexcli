package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]CacheEntry
	mu   *sync.RWMutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		data: make(map[string]CacheEntry),
		mu:   &sync.RWMutex{},
	}
	go c.checkLoop(interval)
	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, found := c.data[key]
	if !found {
		c.data[key] = CacheEntry{
			createdAt: time.Now().UTC(),
			val:       val,
		}
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, found := c.data[key]
	return val.val, found
}

func (c Cache) checkLoop(interval time.Duration) {
	ticket := time.NewTicker(interval)
	for range ticket.C {
		c.check(time.Now().UTC(), interval)
	}
}

func (c Cache) check(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.data, k)
		}
	}
}
