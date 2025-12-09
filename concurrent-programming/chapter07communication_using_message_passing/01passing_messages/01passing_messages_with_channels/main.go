package main

import (
	"fmt"
	"time"
)

func receiver(messages chan string) {
	msg := ""
	for msg != "STOP" {
		msg = <-messages
		fmt.Println("Received:", msg)
	}
}

func NoReceiver(messages chan string) {
	time.Sleep(5 * time.Second) // Sleeps for 5 seconds instead of sending any message
	fmt.Println("Receiver slept for 5 seconds")
}

func main() {
	msgChannel := make(chan string)
	go receiver(msgChannel)
	fmt.Println("Sending HELLO...")
	// Sends three string message over the channel
	msgChannel <- "HELLO"
	fmt.Println("Sending THERE...")
	fmt.Println("Sending STOP...")
	msgChannel <- "STOP"

	fmt.Println("\n== There is no recipient. ==")

	msgChannel2 := make(chan string)
	go NoReceiver(msgChannel2)
	fmt.Println("Reading message from channel...")
	msg := <-msgChannel2 // reads a message from channel
	fmt.Println("Received:", msg)

}
