package creatingachannelwithsemaphores

import (
	"container/list"
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

type Channel[M any] struct {
	capacitySema *Semaphore // capacity semaphore to block sender when the buffer is full
	sizeSema     *Semaphore // buffer size semaphore to block the receiver when the buffer is empty
	mutex        sync.Mutex // mutex protecting our shared list data structure
	buffer       *list.List // linked list to be used as a queue data structure
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		capacitySema: NewSemaphore(capacity), // creates a new semaphore with the number of permits equal to the input capacity
		sizeSema:     NewSemaphore(0),        // creates a new semaphore with the number of permits equal to 0
		buffer:       list.New(),             // creates a new , empty linked list
	}
}
