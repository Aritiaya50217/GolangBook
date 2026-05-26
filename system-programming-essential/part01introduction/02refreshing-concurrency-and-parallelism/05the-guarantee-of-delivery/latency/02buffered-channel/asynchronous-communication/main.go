package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "A"
	ch <- "B"

	fmt.Println("Sent successfully")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
