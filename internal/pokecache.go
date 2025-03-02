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

func NewCache(interval time.Duration) *Cache {
	tick := time.NewTicker(interval)
	nc := &Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go nc.ReapLoop(tick)
	return nc
}

func (ch *Cache) Add(key string, val []byte) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	if _, exists := ch.cache[key]; exists {
		return
	}
	ch.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (ch *Cache) Get(key string) ([]byte, bool) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	v, exists := ch.cache[key]
	if !exists {
		return []byte{}, false
	}
	return v.val, true
}

func (ch *Cache) ReapLoop(interval <-chan *time.Ticker) error {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	for {
		select {
		case _, ok := <-interval:
			if !ok {

			}
		}
	}
}
