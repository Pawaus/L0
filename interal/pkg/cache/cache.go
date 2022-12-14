package cache

import "api/interal/pkg/domain"

type Cache struct {
	data map[string]domain.Order
}

func New() *Cache {
	cache := &Cache{data: make(map[string]domain.Order)}
	return cache
}

func (c *Cache) Get(index string) domain.Order {
	return c.data[index]
}

func (c *Cache) Update(index string, value domain.Order) {
	c.data[index] = value
}
