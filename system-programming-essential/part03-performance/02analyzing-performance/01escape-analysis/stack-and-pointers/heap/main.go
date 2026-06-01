package main

import "fmt"

type person struct {
	name string
	age  int
}

func createPerson() *person {
	p := person{
		name: "Alex Rios",
		age:  99,
	}
	return &p
}

func main() {
	heap1 := createPerson()
	heap2 := createPerson()

	fmt.Println("heap1 : ", heap1)
	fmt.Println("heap2 : ", heap2)
}
