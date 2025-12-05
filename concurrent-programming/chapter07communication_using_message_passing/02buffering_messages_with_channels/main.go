package main

import (
	"fmt"
	"sync"
	"time"
)

func receiver(messages chan int, wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 { // keeps reading messages from the channel until it receives a -1
		time.Sleep(1 * time.Second) // waits for 1 second
		msg = <-messages            // Reads the next message from the channel
		fmt.Println("Received : ", msg)
	}
	wGroup.Done() // calls Done() on the waitgroup after reading all the message
}

func main() {
	msgChannel := make(chan int, 3) // creates a new channel with a buffer capacity of 3 messages
	wGroup := sync.WaitGroup{}      // creates a waitgroup with a size of 1
	wGroup.Add(1)
	go receiver(msgChannel, &wGroup) // start the receiver goroutine with the buffered channel and waitgroup
	for i := 1; i <= 6; i++ {
		size := len(msgChannel)
		fmt.Printf("%s Sending: %d. Buffer Size: %d\n", time.Now().Format("15:04:05"), i, size)
		msgChannel <- i
	}
	msgChannel <- -1  // sends a message containing -1
	wGroup.Wait() // waits on the waitgroup until the receiver is finished

}
