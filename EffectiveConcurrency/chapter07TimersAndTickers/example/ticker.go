package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	fmt.Println("Starting ticker...")
	for {
		select {
		case <-done:
			ticker.Stop()
			fmt.Println("Ticker stopped")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at ", t)
		}
	}
}
