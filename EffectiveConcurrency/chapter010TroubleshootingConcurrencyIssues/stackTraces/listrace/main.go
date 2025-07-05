package main

import (
	"container/list"
	"math/rand"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ll := list.New()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			ll.PushBack(rand.Int())
		}
	}()

	// goroutine that empties the list
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			if ll.Len() > 0 {
				ll.Remove(ll.Front())
			}
		}
	}()
	wg.Wait()
}
