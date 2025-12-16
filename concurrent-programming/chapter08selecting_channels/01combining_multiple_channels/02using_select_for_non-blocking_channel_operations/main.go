package main

import (
	"fmt"
	"time"
)

func sendMsgAfter(seconds time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		time.Sleep(seconds)
		messages <- "Hello"
	}()
	return messages
}

func main() {
	messages := sendMsgAfter(3 * time.Second) // sends channel message after 3 seconds
	for {
		select {
		case msg := <-messages: // reads a message from the channel if there is one
			fmt.Println("Message received: ", msg)
			return // when a message is available, terminates the execution
		default: // when no message is available , the default case is executed.
			fmt.Println("No Messages waiting")
			time.Sleep(1 * time.Second)
		}
	}
}
