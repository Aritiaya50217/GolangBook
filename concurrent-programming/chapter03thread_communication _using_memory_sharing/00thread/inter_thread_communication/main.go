package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("== Shared Memory (Mutex) ==")

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Println("Goroutine", id, "updated counter to", counter)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Final Counter: ", counter)

	fmt.Println("== Channel (Message Passing) ==")
	ch := make(chan string)
	for i := 0; i < 3; i++ {
		go func() {
			index := strconv.Itoa(i)
			ch <- "Hello from Goroutine " + index
		}()
		message := <-ch
		fmt.Println("message : ", message)
	}

}
