package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	var a int
	// Lock mutex in main goroutine
	m.Lock()
	done := make(chan struct{})

	// G1
	go func() {
		// This will block until G2 unlocks mutex
		m.Lock()
		// a = 1 happend-before , so this prints 1
		fmt.Println("a : ", a)
		m.Unlock()
		close(done)
	}()

	// G2
	go func() {
		a = 1
		// G1 will block until this runs
		m.Unlock()
	}()
	<-done

}
