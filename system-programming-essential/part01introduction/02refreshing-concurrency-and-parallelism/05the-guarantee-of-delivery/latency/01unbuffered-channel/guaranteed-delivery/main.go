package main

import "fmt"

func main() {
	// unbuffered channel
	ch := make(chan string)

	go func() {
		msg := <-ch
		fmt.Println("Received : ", msg)
	}()

	ch <- "Hello"

}
