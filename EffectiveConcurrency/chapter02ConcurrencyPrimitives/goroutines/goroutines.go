package main

import (
	"fmt"
	"time"
)

func main() {
	for _, s := range []string{"a", "b", "c"} {
		go func() {
			fmt.Printf("Goroutine %s\n", s)
		}()
	}
	time.Sleep(100)

	// var s string
	// s = "a"
	// go func() {
	// 	fmt.Printf("Goroutine %s\n", s)
	// }()

	// s = "b"
	// go func() {
	// 	fmt.Printf("Goroutine %s\n", s)
	// }()

	// time.Sleep(100)
	// fmt.Println("exist")
}
