package cache

import (
	"sync"
)

// Cache представляет собой простой кеш
type Cache struct {
	mu     sync.Mutex
	values map[string]string
}

// NewCache создает новый экземпляр кеша
func NewCache() *Cache {
	return &Cache{
		values: make(map[string]string),
	}
}

// Add добавляет запись в кеш
func (c *Cache) Add(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

// Delete удаляет запись из кеша
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.values, key)
}

// Get возвращает значение записи по ключу из кеша
func (c *Cache) Get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.values[key]
}
