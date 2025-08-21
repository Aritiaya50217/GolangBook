package main

import "fmt"

func main() {
	goto customerLabel

	// Will never get executed bacause
	// the goto. statement will jump right
	// past this line
	fmt.Println("Hello")
	
customerLabel:
	fmt.Println("World")
}
