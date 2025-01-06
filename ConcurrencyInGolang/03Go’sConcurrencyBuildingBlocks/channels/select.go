package main

import (
	"fmt"
	"time"
)

func single() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read ...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func multiple() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func timeAfter() {
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second): // รอไม่เกิน 1 วินาที
		fmt.Println("Timed out.")
	}
}

func timeSince() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

func timeSleep() {
	done := make(chan interface{})
	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
		fmt.Printf("Achieved %v cycles\n", workCounter)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

func timeAfterAndTimeSleep() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "finished after sleep"
	}()

	select {
	case msg := <-ch:
		fmt.Println("msg : ", msg)
	case <-time.After(1 * time.Second): // รอไม่เกิน 1 วินาที
		fmt.Println("Timeout")
	}
}

func main() {
	single()
	multiple()
	timeAfter()
	timeSince()
	timeSleep()
	timeAfterAndTimeSleep()
}
