package main

import (
	"fmt"
	"time"
)

func main() {
	// create a ticker that ticks every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// create a timer that fires after 10 seconds
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	// use a select statement to handle the signals from ticker and timer
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at ", tick)
		case <-timer.C:
			fmt.Println("Timer expired")
			return

		}
	}
}
