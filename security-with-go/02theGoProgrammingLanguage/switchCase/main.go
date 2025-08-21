package main

import "fmt"

func main() {

	x := 42
	switch x {
	case 25:
		fmt.Println("x is 25")
	case 42:
		fmt.Println("x is 42")
		fallthrough 
	case 100:
		fmt.Println("x is 100")
	default:
		fmt.Println("x is something else.")
	}
}
