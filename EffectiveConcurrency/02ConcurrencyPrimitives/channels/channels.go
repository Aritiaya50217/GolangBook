package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch2 <- 1
	}()

	go func() {
		fmt.Println(<-ch1)
	}()

	select {
	case ch1 <- <-ch2:
		time.Sleep(time.Second)
	default:
	}
}
