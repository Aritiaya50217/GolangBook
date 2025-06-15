package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method
func (p *Person) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  27,
	}
	p.Greet()
}
