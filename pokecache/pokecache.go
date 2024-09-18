package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]CacheEntry
	Mu   *sync.RWMutex
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		Data: make(map[string]CacheEntry),
		Mu:   &sync.RWMutex{},
	}
	go c.checkLoop(interval)
	return &c
}

func (c Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	_, found := c.Data[key]
	if !found {
		c.Data[key] = CacheEntry{
			CreatedAt: time.Now().UTC(),
			Val:       val,
		}
	}
}

// TODO: change the key to url.URL in order to check all the params and the path and host (no matter the order of params)
func (c Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	// query := key.Query()
	val, found := c.Data[key]
	return val.Val, found
}

func (c Cache) checkLoop(interval time.Duration) {
	ticket := time.NewTicker(interval)
	for range ticket.C {
		c.check(time.Now().UTC(), interval)
	}
}

func (c Cache) check(now time.Time, last time.Duration) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for k, v := range c.Data {
		if v.CreatedAt.Before(now.Add(-last)) {
			delete(c.Data, k)
		}
	}
}
