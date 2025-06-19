package pokecache

import (
// 	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	Mu sync.Mutex
	Entries map[string]cacheEntry
	Interval time.Duration
	tkr *time.Ticker
}

func NewCache(interval time.Duration) Cache {
	cache := Cache {
		Interval: interval,
	}
	cache.Entries = make(map[string]cacheEntry)
	cache.tkr = time.NewTicker(interval)

	go cache.reapLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	entry := cacheEntry {
		createdAt: time.Now(),
		val: val,
	}
	c.Entries[key] = entry
	c.Mu.Unlock()
}

func (c Cache) Get(key string) (v []byte, found bool) {
	c.Mu.Lock()
	entry, found := c.Entries[key]
	c.Mu.Unlock()
	if !found {
		return nil, found
	} else {
		return entry.val, found
	}
}

func (c Cache) reapLoop() {
	for {
		select {
		case <-c.tkr.C:
			c.Mu.Lock()
			for k,v := range c.Entries {
				if time.Since(v.createdAt) > c.Interval {
					delete(c.Entries, k)
				}
			}
			c.Mu.Unlock()
		}
	}
}


