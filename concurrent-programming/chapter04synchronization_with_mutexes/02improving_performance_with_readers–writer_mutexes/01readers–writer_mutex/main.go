package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchRecorder(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		// Protects critical section with a write mutex
		*matchEvents = append(*matchEvents, "Match event "+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Append match event")
	}
}

func clientHandler(mEvents *[]string, mutex *sync.RWMutex, st time.Time) {
	for i := 0; i < 100; i++ {
		mutex.Lock()
		// Protects critical section with a read mutex
		allEvents := copyAllEvents(mEvents)
		mutex.Unlock()

		// Calculates the time taken since the start
		timeTaken := time.Since(st)
		// Outputs to the console the time taken to serve the client
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, 0, len(*matchEvents))
	for _, e := range *matchEvents {
		allEvents = append(allEvents, e)
	}
	return allEvents
}

func main() {
	mutex := sync.RWMutex{}
	var matchEvents = make([]string, 0, 10000)
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "Match event")
	}
	// Passes readers–writer mutex to match recorder
	go matchRecorder(&matchEvents, &mutex)
	// records the start time before starting the client handler goroutines
	start := time.Now()
	for j := 0; j < 5000; j++ {
		// Passes readers–writer mutex to client handler goroutine
		go clientHandler(&matchEvents, &mutex, start)
	}
	time.Sleep(100 * time.Second)
}
