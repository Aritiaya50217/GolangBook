package main

import (
	"log"
	"time"
)

func countDown() {
	for i := 5; i >= 0; i-- {
		log.Println(i)
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	// Kick off a thread
	go countDown()

	// Since functions are first-class
	// you can write an anonymous function
	// for a goroutine
	go func() {
		time.Sleep(time.Second * 2)
		log.Println("Deleyed greetings!")
	}()

	// Use channel to signal when complete
	// or in the case just wait
	time.Sleep(time.Second * 4)
}
