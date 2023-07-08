package cache

import (
	"container/list"
	"sync"
)

// Cache item
type Item struct {
	key   string
	value any
}

// Custom LRU Cacche implementation
type LRUCache struct {
	items    map[string]*list.Element
	queue    *list.List
	capacity int
	dataMu   sync.RWMutex
}

// Initialize new lru cache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		items:    make(map[string]*list.Element),
		queue:    list.New(),
		capacity: capacity,
	}
}

// Get element from cache by key
func (c *LRUCache) Get(key string) (any, bool) {
	c.dataMu.RLock()
	defer c.dataMu.RUnlock()

	if elem, exists := c.items[key]; exists {
		c.queue.MoveToFront(elem)
		return elem.Value.(*Item).value, true
	}

	return nil, false
}

// Save elemnt in cache
func (c *LRUCache) Set(key string, val any) error {
	c.dataMu.Lock()
	defer c.dataMu.Unlock()

	// Проверяем, существует ли элемент с заданным ключем
	if elem, exists := c.items[key]; exists {
		elem.Value.(*Item).value = val
		c.queue.MoveToFront(elem)
		return nil
	}

	// Если элемента нет, создаем новый и добавляем в кэш
	newItem := &Item{key: key, value: val}
	elem := c.queue.PushFront(newItem)
	c.items[key] = elem

	// Если количество превышает вместимость кэша, удаляем самый старый
	if c.queue.Len() > c.capacity {
		oldestElem := c.queue.Back()
		if oldestElem != nil {
			oldestElem := c.queue.Remove(oldestElem).(*Item)
			delete(c.items, oldestElem.key)
		}
	}

	return nil
}

// Delete element from cache by key
func (c *LRUCache) Delete(key string) {
	c.dataMu.Lock()
	defer c.dataMu.Unlock()

	if elem, exists := c.items[key]; exists {
		delete(c.items, key)
		c.queue.Remove(elem)
	}
}

// Return count of elements in cache
func (c *LRUCache) Count() int {
	c.dataMu.Lock()
	defer c.dataMu.Unlock()

	return c.queue.Len()
}

// Clear cache
func (c *LRUCache) Clear() {
	c.dataMu.Lock()
	defer c.dataMu.Unlock()

	c.items = make(map[string]*list.Element)
	c.queue.Init()
}
