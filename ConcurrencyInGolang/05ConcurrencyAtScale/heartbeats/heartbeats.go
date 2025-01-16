package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}) (<-chan interface{}, <-chan int) {
		heartbeatStraem := make(chan interface{}, 1)
		workStream := make(chan int)

		go func() {
			defer close(heartbeatStraem)
			defer close(workStream)

			time.Sleep(2 * time.Second)

			for i := 0; i < 10; i++ {
				select {
				case heartbeatStraem <- struct{}{}:
				default:
				}

				select {
				case <-done:
					return
				case workStream <- rand.Intn(10):
				}
			}
		}()

		return heartbeatStraem, workStream
	}
	_ = doWork

	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}
	}
}
