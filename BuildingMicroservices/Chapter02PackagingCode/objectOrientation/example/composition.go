package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("My name is ", a.Name)
}

type Dogs struct {
	Animal // embedded struct (composition)
	Breed  string
}

func main() {
	d := Dogs{Animal: Animal{Name: "Bobby"}, Breed: "Beagle"}
	d.Speak()
}
