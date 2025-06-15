package main

import "fmt"

type Animal struct {
	Name   string
	CanFly bool
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Grow() {
	p.Age++
}

func (p Person) DoesNotGrow() {
	p.Age++
}

func main() {
	anAnimal := Animal{Name: "Lion", CanFly: false}
	fmt.Println(anAnimal)

	p := Person{"JY", 10}
	p.Grow()
	fmt.Println(p.Age)

	ptr := &p
	ptr.DoesNotGrow()
	fmt.Printf("from p : %v , from prt : %v\n", p.Age, ptr.Age)
}
