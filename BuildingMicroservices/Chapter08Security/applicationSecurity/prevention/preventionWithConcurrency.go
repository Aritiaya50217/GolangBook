package main

import (
	"fmt"
	"sync"
)

// ป้องกัน race condition หรือ deadlock
func main() {
	var mu sync.Mutex
	count := 0
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Count : ", count)
}
