package main

import (
	"fmt"
	"sync"
	"time"
)

type ReadWriteMutex struct {
	readersCounter int
	readersLock    sync.Mutex // mutex for synchronizing readers access
	globalLock     sync.Mutex // mutex for blocking any writers access
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock() // synchronizes access so that only one goroutine is allowed at any time
	rw.readersCounter++   // reader goroutine increments readersCounter by 1
	if rw.readersCounter == 1 {
		rw.globalLock.Lock() // if a reader goroutine is the first one in , it attempts to lock globalLock
	}
	rw.readersLock.Unlock() // synchronizes access so that only one goroutine is allowed at any time
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock() // any writer access requires a lock on globalLock
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter-- // the reader goroutine decrements readersCounter by 1
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock() // it the reader goroutine is the last one out,it unlocks the global lock
	}
	rw.readersLock.Unlock() // synchronizes access so that only one goroitine is allowed at any time
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock() // the writer goroutine, finishing its critical section , releases the global lock.
}

func main() {
	rwMutex := ReadWriteMutex{}
	for i := 0; i < 10; i++ {
		go func() {
			rwMutex.ReadLock()
			fmt.Println("Read started")
			time.Sleep(5 * time.Second)
			fmt.Println("Read done")
			rwMutex.ReadUnlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Write started")
	rwMutex.WriteLock()
	fmt.Println("Write finished")
}
