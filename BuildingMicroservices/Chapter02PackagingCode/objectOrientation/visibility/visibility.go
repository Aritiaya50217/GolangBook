package main

import "fmt"

type Person struct {
	Name string // public
	age  int    //private
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) setAge(a int) {
	p.age = a
}

func main() {
	person := Person{
		age: 27,
	}
	age := person.GetAge()
	fmt.Println("age : ", age)

	person.setAge(20)
	fmt.Println("age : ", person.age)

}
