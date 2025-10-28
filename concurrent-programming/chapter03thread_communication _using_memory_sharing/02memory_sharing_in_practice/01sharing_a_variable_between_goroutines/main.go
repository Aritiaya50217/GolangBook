package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		*seconds -= 1 // the goroutine updates the value go the shared variable.
	}
}

func main() {

	count := 5
	go countdown(&count)

	// the main() goroutine reads the value of the shared variable every half second.
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}

	// ใช้ mutex
	fmt.Println("== sync.Mutex ==")
	var mu sync.Mutex
	var sum int

	for i := 0; i < 5; i++ {
		go func() {
			mu.Lock()
			sum++
			mu.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Summary : ", sum)

}
