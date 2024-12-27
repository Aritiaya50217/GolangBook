package main

import (
	"fmt"
	"sync"
)

func Once() {
	var count int
	increment := func() {
		count++
	}
	var once sync.Once
	once.Do(increment)

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Println("Count is ", count)
}

func main() {
	Once()

	var onceA, onceB sync.Once
	var initB func()
	initA := func() {
		onceB.Do(initB)
	}
	initB = func() {
		onceA.Do(initA)
	}
	onceA.Do(initA)
}
