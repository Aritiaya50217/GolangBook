package main

import (
	"fmt"
	"time"
)

func safeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		fn()
	}()
}

func main() {
	safeGo(func() {
		fmt.Println("Starting goroutine...")
		panic("something went wrong") // ทำให้เกิด panic
	})

	time.Sleep(1 * time.Second)
	fmt.Println("Main function continues running.")
}
