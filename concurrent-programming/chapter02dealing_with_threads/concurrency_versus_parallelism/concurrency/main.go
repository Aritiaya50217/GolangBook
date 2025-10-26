package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s:  %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go task("Task A")
	go task("Task B")

	time.Sleep(3 * time.Second)
	fmt.Println("All done!")
}
