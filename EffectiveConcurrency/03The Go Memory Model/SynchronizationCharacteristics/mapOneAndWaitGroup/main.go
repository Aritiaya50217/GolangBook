package main

import (
	"fmt"
	"sync"
)

type Data struct {
	ID      string
	Payload string
}

type Cache struct {
	values sync.Map
}

type cachedValue struct {
	sync.Once
	value *Data
}

func loadData(id string) *Data {
	return &Data{ID: id, Payload: "test"}
}

func (c *Cache) Get(id string) *Data {
	// Get the cached value , or store an empty value
	v, _ := c.values.LoadOrStore(id, &cachedValue{})
	cv := v.(*cachedValue)
	// if not initialized , initialize here
	cv.Do(func() {
		cv.value = loadData(id)
	})
	return cv.value
}

func main() {
	var data *Data
	var c Cache

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		data = c.Get("1")
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("id : %s , payload : %s", data.ID, data.Payload)

}
