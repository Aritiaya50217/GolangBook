package main

import (
	"fmt"
	"sync"
)

type WaitGrp struct {
	groupSize int        // the waitgroup size property,initialized to 0 by default.
	cond      *sync.Cond // the condition variable to be used in the waitgroup.
}

func NewWaitGrp() *WaitGrp {
	return &WaitGrp{
		cond: sync.NewCond(&sync.Mutex{}), // initializes the condition variable with a new mutex.
	}
}

func (wg *WaitGrp) Add(delta int) {
	wg.cond.L.Lock()      // protects the update to groupSize with a mutex lock on the condition variable.
	wg.groupSize += delta // increases groupSize by delta
	wg.cond.L.Unlock()
}

func (wg *WaitGrp) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		wg.cond.Wait() // waits and atomically releases the mutex while groupSize is greater than 0.
	}
	wg.cond.L.Unlock()
}

func (wg *WaitGrp) Done() {
	wg.cond.L.Lock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast() // if it's the last goroutine to be done in the waitgroup , it broadcasts on the condition variable.
	}
	wg.cond.L.Unlock()
}

func doWork(id int, wg *WaitGrp) {
	fmt.Println(id, "Done working")
	wg.Done()
}

func main() {
	wg := NewWaitGrp()
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go doWork(i, wg)
		go doWork(i, wg)
	}
	wg.Wait()
	fmt.Println("All complate")
}
