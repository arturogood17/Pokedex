package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	nc := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go nc.reapLoop(interval)
	return nc
}

func (ch *Cache) Add(key string, value []byte) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	ch.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (ch *Cache) Get(key string) ([]byte, bool) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	v, exists := ch.cache[key]
	return v.val, exists
}

func (ch *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	for range tick.C {
		ch.reap(time.Now().UTC(), interval)
	}
}

func (ch *Cache) reap(now time.Time, interval time.Duration) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	for k, v := range ch.cache {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(ch.cache, k)
		}
	}
}
