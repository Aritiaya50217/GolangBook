package main

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	permits int        // Permits remaining on the semaphore
	cond    *sync.Cond // condition variable used for waiting when there are not enough permits
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		permits: n,                           // initial number of permits on the new semaphore.
		cond:    sync.NewCond(&sync.Mutex{}), // initializes a new condition variable and associated mutex on the new semaphore.
	}
}

func (rw *Semaphore) Acquire() {
	rw.cond.L.Lock() // acquires mutex to protect permits variable.
	for rw.permits <= 0 {
		rw.cond.Wait() // waits until there is an available permit.
		fmt.Println("Wait")
	}
	rw.permits--       // decreases the number of available permits by 1.
	rw.cond.L.Unlock() // releases mutex
}

func (rw *Semaphore) Release() {
	rw.cond.L.Lock()
	rw.permits++     // increases the number of valiable permits by 1
	rw.cond.Signal() // signals condition variable that one more permit is available.
	rw.cond.L.Unlock()
}

func doWork(semaphore *Semaphore) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	semaphore.Release()
}

func main() {
	semaphore := NewSemaphore(0)
	for i := 0; i < 5; i++ {
		go doWork(semaphore)
		fmt.Println("Waiting for child goroutine")
		semaphore.Acquire()
		fmt.Println("Child goroutine finished.\n")
	}
}
