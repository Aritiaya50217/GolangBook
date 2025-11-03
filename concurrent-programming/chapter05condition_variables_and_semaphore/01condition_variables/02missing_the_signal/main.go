package main

import (
	"fmt"
	"sync"
)

func doWork(cond *sync.Cond) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	cond.L.Lock()   // Locks mutex before signaling
	cond.Signal()   // Signals on condition variable
	cond.L.Unlock() // Unlocks mutex after signaling
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	// repeats 5000 times
	for i := 0; i < 5000; i++ {
		fmt.Println("index : ", i)
		// Starts a goroutine , simulating doing some work
		go doWork(cond)
		cond.Wait() // waits for a finished signal from the goroutine
		fmt.Println("Child goroutine finished")
	}
	cond.L.Unlock()
}
