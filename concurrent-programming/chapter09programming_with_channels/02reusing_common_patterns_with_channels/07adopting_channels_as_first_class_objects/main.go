package main

import "fmt"

func primeMultipleFilter(numbers <-chan int, quit chan<- int) {
	var right chan int
	p := <-numbers
	fmt.Println(p)           // receives the first message containing the prime number p on the input channel and prints it
	for n := range numbers { // reads the next number from the input channel
		if n%p != 0 { // discards any received number that is a multiple of p
			if right == nil {
				right = make(chan int)
				go primeMultipleFilter(right, quit) // if the current goroutine has no right , is starts a new goroutine and connects to it with a channel.
			}
			right <- n // Passes the filtered number to the right channel
		}
	}
	if right == nil {
		close(quit) // closes the quit channel if there are no more numbers to filter and the goroutine has right channel
	} else {
		close(right) // otherwise , closes right channel
	}
}

func main() {
	numbers := make(chan int)             // creates an input channel that will feed the prime multiple filters
	quit := make(chan int)                // creates a common quit channel
	go primeMultipleFilter(numbers, quit) // starts the first goroutine in the pipline , passing the numbers and quit channels
	for i := 2; i < 100000; i++ { // feeds sequential numbers, starting from 2 up to 100,000 onto the input channel 
		numbers <- i
	}
	close(numbers) // closes the input channel , signaling that there will be no more numbers
	<-quit // waits for the quit channel close
}
