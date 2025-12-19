package main

import (
	"fmt"
	"time"
)

const (
	passwordToGuess = "go far"                      // sets the password that we need to guess
	alphabet        = " abcdefghijklmnopqrstuvwxyz" // Defines all possible characters that the password is made of.
)

func toBase27(n int) string {
	result := ""
	for n > 0 {
		result = string(alphabet[n%27]) + result // algorithm converts a decimal integer into a string of base 27 using the alpgabet constant.
		n /= 27
	}
	return result
}

func guessPassword(from int, upto int, stop chan int, result chan string) {
	for guessN := from; guessN < upto; guessN += 1 {
		select {
		case <-stop:
			fmt.Printf("Stopped at %d [%d,%d]\n", guessN, from, upto)
			return

		default:
			if toBase27(guessN) == passwordToGuess { // checks whether the password matches (in a real-life system , we would try to acces the protected resource)
				result <- toBase27(guessN) // sends matching password on the result channel.
				close(stop)                // closes the channel so that other goroutines stop checking the password
				return
			}
		}
	}
	fmt.Printf("Not found between [%d,%d)\n", from, upto)
}

func main() {
	finished := make(chan int)
	passwordFound := make(chan string) // creates a channel that will contain the discovered password after it's found

	for i := 1; i <= 387_420_488; i += 10_000_000 {
		go guessPassword(i, i+10_000_000, finished, passwordFound) // Creates a goroutine with input ranges[1, 10M), [10M, 20M), . . . [380M, 390M)
	}
	fmt.Println("password found:", <-passwordFound) // Waits for the password to be found
	close(passwordFound)
	time.Sleep(5 * time.Second)
}
