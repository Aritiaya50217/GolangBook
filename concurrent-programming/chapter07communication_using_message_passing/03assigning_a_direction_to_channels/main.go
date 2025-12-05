package main

import (
	"fmt"
	"time"
)

func receiver(message <-chan int) { // Declares a receiv-only channel
	fmt.Println("start ...  : ", <-message)
	for {
		msg := <-message // receives messages from the channel
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
	}
}

func sender(messages chan<- int) { // Declares a send-only channel
	for i := 1; ; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending:", i)
		messages <- i
		time.Sleep(1 * time.Second) // sends a message on the channel every second
	}

}

func main() {
	msgChannel := make(chan int)
	go receiver(msgChannel)
	go sender(msgChannel)
	time.Sleep(5 * time.Second)
}
