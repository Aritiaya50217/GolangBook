package main

import (
	"fmt"
)

func printNumbers(numbers <-chan int, quit chan int) {
	go func() {
		for i := 0; i < 10; i++ { // consumes 10 items from the numbers channel
			fmt.Println(<-numbers)
		}
		close(quit) // <- close the quit channel
	}()
}

func main() {
	numbers := make(chan int) // creates the numbers and quit channels
	quit := make(chan int)
	printNumbers(numbers, quit) // call the printNumbers() function , passing the channel
	next := 0
	for i := 1; ; i++ {
		next += i // generates the next triangular number
		select {
		case numbers <- next: // sends the number on the numbers channel
		case <-quit:
			fmt.Println("Quitting number generation") // when the quit channel is unblocked , outputs a message and terminates the execution
			return
		}
	}
}
