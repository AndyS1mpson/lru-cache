package cache

import (
	"container/list"
	"sync"
)

type Item struct {
	Val       any
	callCount int
}

type LRUCache struct {
	items map[string]*list.Element
	queue    *list.List
	capacity int
	dataMu   sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		items: make(map[string]*list.Element),
		queue: list.New(),
		capacity: capacity,
	}
}

func (c *LRUCache) Get(key string) (any, bool)

func (c *LRUCache) Set(key string, val any) error

func (c *LRUCache) Delete(key string)

func (c *LRUCache) Count() int

func (c *LRUCache) Clear()
