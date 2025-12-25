package main

import (
	"fmt"
	"math"
	"math/rand"
)

// accepts numbers in the inputs channel and returns a channel containing only prime numbers
func primaryOnly(inputs <-chan int) <-chan int {
	results := make(chan int)
	go func() {
		for c := range inputs {
			isPrime := c != 1 // checks to ensure c is not 1 , since 1 is not a prime
			for i := 2; i <= int(math.Sqrt(float64(c))); i++ {
				if c%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				results <- c // if c is prime , outputs c on the results channel
			}
		}
	}()
	return results
}

func main() {
	numbersChannel := make(chan int)
	primes := primaryOnly(numbersChannel)
	for i := 0; i < 100; {
		select {
		case numbersChannel <- rand.Intn(1000000000) + 1: // feeds a random number between 1 and 1 billopn onto the input isPrimeChannel
		case p := <-primes: // reads an output prime number
			fmt.Println("Found prime : ", p)
			i++
		}
	}
}
