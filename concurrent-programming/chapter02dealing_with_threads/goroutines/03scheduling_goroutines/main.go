package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("Goroutine %d on thread %d (iteration %d)\n", id, runtime.NumGoroutine(), i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("== sayHello ==")
	go sayHello()
	// Calling the Go scheduler gives the other goroutine a chance to run.
	runtime.Gosched()
	fmt.Println("Finished")

	fmt.Println("== worker ==")
	runtime.GOMAXPROCS(2) // ใช้ 2 core เท่านั้น
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutines done")

}
