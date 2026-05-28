package main

import "fmt"

func createUser() *string {
	name := "John"
	return &name
}

func main() {
	name := createUser()
	fmt.Printf("address in memory : %v\ninformation allocated on the heap is : %v\n", name, *name)
}
