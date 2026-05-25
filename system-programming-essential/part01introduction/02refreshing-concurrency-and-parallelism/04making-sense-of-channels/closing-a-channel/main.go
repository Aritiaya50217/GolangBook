package main

import (
	"fmt"
	"sync"
)

func throwBalls(color string, balls chan string) {
	balls <- color
}

func main() {
	balls := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		throwBalls("red", balls)
	}()

	go func() {
		defer wg.Done()
		throwBalls("green", balls)
	}()

	go func() {
		wg.Wait()
		close(balls)
	}()
	
	for color := range balls {
		fmt.Println(color)
	}

}
