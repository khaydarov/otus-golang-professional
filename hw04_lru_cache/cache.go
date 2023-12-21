package hw04lrucache

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
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if v, ok := c.items[key]; ok {
		v.Value = value
		c.queue.MoveToFront(v)

		return true
	}

	if c.queue.Len() == c.capacity {
		b := c.queue.Back()
		c.queue.Remove(b)

		for k, v := range c.items {
			if b == v {
				delete(c.items, k)
			}
		}
	}

	i := c.queue.PushFront(value)
	c.items[key] = i

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if v, ok := c.items[key]; ok {
		c.queue.MoveToFront(v)
		return v.Value, true
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
