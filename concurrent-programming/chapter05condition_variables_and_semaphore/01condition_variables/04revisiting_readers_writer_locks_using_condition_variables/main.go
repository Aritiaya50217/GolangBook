package main

import (
	"fmt"
	"sync"
	"time"
)

type ReadWriteMutex struct {
	readersCounter int  // stores the number of readers currently holding the read lock
	writersWaiting int  // stores the number of writers currently waiting
	writerActive   bool // indicates if a writer is holding the write lock
	cond           *sync.Cond
}

func NewReadWriteMutex() *ReadWriteMutex {
	return &ReadWriteMutex{
		cond: sync.NewCond(&sync.Mutex{}), // initializes a new ReadWriteMutex with a new condition variable and associated mutex
	}
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.cond.L.Lock()                               // acquires mutex
	for rw.writersWaiting > 0 || rw.writerActive { // waits on condition variable while writers are waiting or active
		rw.cond.Wait()
	}
	rw.readersCounter++ // increments readers'counter
	rw.cond.L.Unlock()  // releases mutex
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.cond.L.Lock()    // acquires mutex
	rw.writersWaiting++ // increments the writers' waiting counter
	for rw.readersCounter > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.writersWaiting--    // once the wait is over , decrements the writers' waiting counter.
	rw.writerActive = true // once the wait is over, marks writer active floag.

	rw.cond.L.Unlock() // releases mutex
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.cond.L.Lock()    // acquires mutex
	rw.readersCounter-- // decrement readers' counter by 1
	if rw.readersCounter == 0 {
		rw.cond.Broadcast() // sends broadcast if the goroutine is the last remaining reader.
	}
	rw.cond.L.Unlock() // releases mutex
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.cond.L.Lock()        // acquires mutex
	rw.writerActive = false // unmarks writer active flag.
	rw.cond.Broadcast()     // sends a broadcast
	rw.cond.L.Unlock()      // releases mutex
}

func main() {
	rwMutex := NewReadWriteMutex()
	for i := 0; i < 2; i++ {
		go func() {
			for {
				rwMutex.ReadLock()
				time.Sleep(1 * time.Second)
				fmt.Println("Read done. ")
				rwMutex.ReadUnlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	rwMutex.WriteLock()
	fmt.Println("Write finished. ")
}
