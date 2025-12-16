package main

import (
	"fmt"
	"time"
)

func writeEvery(msg string, seconds time.Duration) <-chan string {
	message := make(chan string) // Creates a new channel of type string
	go func() {                  //creates a new , anonymous goroutine
		for {
			time.Sleep(seconds) // sleeps for specified period
			message <- msg      // sends the specified message on the channel

		}
	}()
	return message // returns the newly created message channel
}

func main() {
	messageFromA := writeEvery("Tick", 1*time.Second) // creates a goroutine sending messages every second on channel A
	messageFromB := writeEvery("Tock", 3*time.Second) // creates a goroutine sending messages every 3 seconds on channel B
	for {                                             // loops forever
		select {
		case msg1 := <-messageFromA: // outputs message from channel A if one is available
			fmt.Println(msg1)
		case msg2 := <-messageFromB: // outputs message from channel B if one is available
			fmt.Println(msg2)
		}
	}

}
