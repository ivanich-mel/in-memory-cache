package inmemorycache

import (
	"fmt"
)
type Cache struct {
	items map[string]interface{}
}
func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}
func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = value
}
func (c* Cache) Get(key string) (interface{}, error){
	v, ok := c.items[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return v, nil
}
func (c *Cache) Delete(key string) {
	delete(c.items, key)
}