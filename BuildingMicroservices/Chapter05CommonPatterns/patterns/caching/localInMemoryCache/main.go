package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data sync.Map
	ttl  time.Duration
}

type cacheItem struct {
	value      string
	expiration int64
}

func (c *Cache) Set(key, value string) {
	c.data.Store(key, cacheItem{
		value:      value,
		expiration: time.Now().Add(c.ttl).Unix(),
	})
}

func (c *Cache) Get(key string) (string, bool) {
	v, ok := c.data.Load(key)
	if !ok {
		fmt.Println("invalid key")
		return "", false
	}

	item := v.(cacheItem)
	if time.Now().Unix() < item.expiration {
		return item.value, true
	}
	c.data.Delete(key) // expired

	return "", false
}

func main() {
	cache := Cache{ttl: 5 * time.Second}
	cache.Set("user:1", "Alice")

	v, ok := cache.Get("user:1")
	if ok {
		fmt.Println("From Cache : ", v)
	}

	time.Sleep(6 * time.Second)

	if _, ok := cache.Get("user:1"); !ok {
		fmt.Println("Cache expired")
	}
}
