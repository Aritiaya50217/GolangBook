package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("Queued", i)
		}
		close(ch)
	}()

	time.Sleep(2 * time.Second)

	for job := range ch {
		fmt.Println("Processing ", job)
	}
}
