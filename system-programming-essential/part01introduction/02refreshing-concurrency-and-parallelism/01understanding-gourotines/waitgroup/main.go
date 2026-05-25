package main

import (
	"fmt"
	"sync"
)

func sayWaitGroup(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go sayWaitGroup("hello", &wg)
	go sayWaitGroup("world", &wg)

	wg.Wait()
}
