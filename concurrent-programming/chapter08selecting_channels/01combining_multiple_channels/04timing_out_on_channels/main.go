package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// sends the message "Hello" on the returned channel after the specified number of seconds.
func sendMsgAfter(seconds time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		time.Sleep(seconds)
		messages <- "Hello"
	}()
	return messages
}

func main() {
	t, _ := strconv.Atoi(os.Args[1])          // reads the timeout value from the program argument
	messages := sendMsgAfter(3 * time.Second) // starts a goroutine that sends a message on the returned channel after 3 seconds.
	timeoutDuration := time.Duration(t) * time.Second
	select {
	case msg := <-messages: // Reads a message from the messages channel if there is one.
		fmt.Println("Message received:", msg)
	case timeNow := <-time.After(timeoutDuration): // creates a channel and timer , receiving a message after the specified duration.
		fmt.Println("Timed out. Waited until:", timeNow.Format("15:04:05"))
	}
}
