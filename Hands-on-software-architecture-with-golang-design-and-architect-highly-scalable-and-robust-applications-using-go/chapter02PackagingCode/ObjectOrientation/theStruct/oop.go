package main

import "fmt"

type Animal struct {
	Name   string
	canFly bool
	Age    int64
}

type Person struct {
	Name string
	Age  int
}

func (p Person) canVote() bool {
	return p.Age > 18
}

func (p *Person) Grow() {
	p.Age++
}

func (p *Person) DoesNotGrow() {
	p.Age++
}

func main() {
	anAnimal := Animal{
		Name:   "Lion",
		canFly: false,
	}
	fmt.Println("anAnimal : ", anAnimal)

	aLionPtr := &anAnimal
	fmt.Println("aLionPtr : ", aLionPtr.Age)

	p := Person{"JY", 10}
	p.Grow()
	fmt.Println("Age : ", p.Age)
	ptr := &p
	ptr.DoesNotGrow()
	fmt.Println(p.Age)
}
