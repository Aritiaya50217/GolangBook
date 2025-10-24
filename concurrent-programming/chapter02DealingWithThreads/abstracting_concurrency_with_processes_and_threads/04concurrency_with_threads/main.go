package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, "iteration", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go task("Thread A")
	go task("Thread B")

	time.Sleep(time.Second * 2)
	fmt.Println("Threads (goroutines) finished")
}
