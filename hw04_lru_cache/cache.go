package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem

	mu sync.Mutex
}

type CacheItem struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheItem := CacheItem{
		Key:   key,
		Value: value,
	}

	if v, ok := c.items[key]; ok {
		v.Value = cacheItem
		c.queue.MoveToFront(v)

		return true
	}

	if c.queue.Len() == c.capacity {
		b := c.queue.Back()
		c.queue.Remove(b)
		evictingCacheItem := b.Value.(CacheItem)

		delete(c.items, evictingCacheItem.Key)
	}

	i := c.queue.PushFront(cacheItem)
	c.items[key] = i

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.items[key]; ok {
		c.queue.MoveToFront(v)

		cacheItem := v.Value.(CacheItem)
		return cacheItem.Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
