package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	size      int // total number of participants in the barrier
	waitCount int // counter variable representing the number of currently suspened executions.
	cond      *sync.Cond
}

func NewBarrier(size int) *Barrier {
	condVar := sync.NewCond(&sync.Mutex{})     // creates new condition variable.
	return &Barrier{size: size, cond: condVar} // creates and returns reference to new barrier.
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount += 1 // increments the count variable by 1.

	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast() // if waitCount has reached the barrier size, resets waitCount and broadcasts on the condition variable
	} else {
		b.cond.Wait() // If waitCount hasn’t reached the barrier size, waits on the condition variable.
	}
	b.cond.L.Unlock()
}

func workAndWait(name string, timeToWork int, barrier *Barrier) {
	start := time.Now()
	for {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(timeToWork) * time.Second) // simulates doing work for a number of seconds.
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		barrier.Wait() // waits for other goroutines to catch up.
	}
}

func main() {
	barrier := NewBarrier(2)            // creates a new barrier with two participants using the implementation.
	go workAndWait("Red", 4, barrier)   // Starts goroutine with the name Red and a timeToWork of 4
	go workAndWait("Blue", 10, barrier) // Starts goroutine with the name Blue and a timeToWork of 10

	time.Sleep(100 * time.Second) // waits for 100 seconds.
}
