package main

import (
	"fmt"
	"sync"
	"time"
)

func Cound() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFormQueue := func(delay time.Duration, i int) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue : ", i)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 1 {
			c.Wait()
		}
		fmt.Println("Adding to queue : ", i)
		queue = append(queue, struct{}{})
		go removeFormQueue(1*time.Second, i)
		c.L.Unlock()
	}
}

func main() {
	Cound()
}
