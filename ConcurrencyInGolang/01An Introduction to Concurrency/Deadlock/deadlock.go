package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func DeadLock() {
	var wg sync.WaitGroup
	sum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Println("sum : ", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)
	go sum(&a, &b)
	go sum(&b, &a)
	wg.Wait()
}

func main() {
	fmt.Println("DeadLock : ")
	DeadLock()
}
