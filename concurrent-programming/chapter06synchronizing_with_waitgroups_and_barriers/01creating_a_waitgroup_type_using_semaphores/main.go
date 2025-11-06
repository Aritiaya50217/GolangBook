package main

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	permits int        // Permits remaining on the semaphore
	cond    *sync.Cond // condition variable used for waiting when there are not enough permits
}

type WaitGrp struct {
	sema *Semaphore // stores semaphore reference on WaitGrp type
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		permits: n,                           // initial number of permits on the new semaphore.
		cond:    sync.NewCond(&sync.Mutex{}), // initializes a new condition variable and associated mutex on the new semaphore.
	}
}

func NewWaitGrp(size int) *WaitGrp {
	return &WaitGrp{
		sema: NewSemaphore(1 - size), // initializes a new semaphore with 1 - size permits
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

func (wg *WaitGrp) Wait() {
	wg.sema.Acquire() // call Acquire() on the semaphore in the Wait() function
}

func (wg *WaitGrp) Done() {
	wg.sema.Release() // when done , calls Release() on the semaphore
}

func doWork(id int, wg *WaitGrp) {
	fmt.Println(id, "Done working")
	wg.Done() // when the goroutine is complete, it calls Done() on the waitgroup.
}

func main() {
	wg := NewWaitGrp(4)
	for i := 1; i <= 4; i++ {
		go doWork(i, wg) // creates a goroutine, passing a reference to the waitgroup.
	}
	wg.Wait() // waits on the waitgroup for work to be complete.
	fmt.Println("All complete")
}
