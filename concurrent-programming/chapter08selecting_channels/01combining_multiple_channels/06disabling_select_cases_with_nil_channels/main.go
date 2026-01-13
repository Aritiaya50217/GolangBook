package main

import (
	"math/rand"
	"time"
)

func generateAmounts(n int) <-chan int {
	amounts := make(chan int) // creates an output channel
	go func() {
		defer close(amounts) // closes
		for i := 0; i < n; i++ {
			amounts <- rand.Intn(100) + 1
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return amounts // return

}

func main() {
	// 199
}
