package main

import (
	"fmt"
	"time"
)

func say(what string) {
	fmt.Println(what)
}

func main() {
	message := "Hello"
	go say(message)
	time.Sleep(5 * time.Second)
}
