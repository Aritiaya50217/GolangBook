package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateTemp() chan int {
	output := make(chan int)
	go func() {
		temp := 50 // fahrenheit
		for {
			output <- temp
			temp += rand.Intn(3) - 1
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

func outputTemp(input chan int) {
	go func() {
		for {
			fmt.Println("Current temp : ", <-input)
			time.Sleep(2 * time.Second)
		}
	}()
}

func generateNumber() chan int {
	output := make(chan int)
	go func() {
		for {
			output <- rand.Intn(10)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

func player() chan string {
	output := make(chan string)
	count := rand.Intn(100)
	move := []string{"UP", "DOWN", "LEFT", "RIGHT"}
	go func() {
		defer close(output)
		for i := 0; i < count; i++ {
			output <- move[rand.Intn(4)]
			d := time.Duration(rand.Intn(200))
			time.Sleep(d * time.Millisecond)
		}
	}()
	return output
}

func main() {
	temp := generateTemp()
	outputTemp(temp)

	number := generateNumber()
	go func() {
		for {
			fmt.Println("Random number : ", <-number)
		}
	}()

	playerCh := player()
	go func() {
		for move := range playerCh {
			fmt.Println("Player move : ", move)
		}
		fmt.Println("Player finished")
	}()

	select {
	case timeNow := <-time.After(10 * time.Second):
		fmt.Println("Game finished at :", timeNow.Format("15:04:05"))
	}

}
