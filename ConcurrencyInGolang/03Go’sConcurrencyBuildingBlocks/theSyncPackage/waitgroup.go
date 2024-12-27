package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()
	wg.Wait()
	fmt.Println("All goroutines complete")
}
func main() {
	WaitGroup()
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Println("Hello from : ", id)
	}
	const numGreets = 5
	var wg sync.WaitGroup
	wg.Add(numGreets)
	for i := 0; i < numGreets; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
