package main

import "fmt"

type Person struct {
	Name string
}

// Person function receiver
func (p Person) PrintInfo() {
	fmt.Println(p.Name)
}

// Person pointer receiver
// If you did not use the pointer receivers
// it would not modify the person object
// Try removing the asterisk here and seeing how the
// program changes behavior

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func main() {
	nanodano := Person{Name: "NanoDano"}
	nanodano.PrintInfo()
	nanodano.ChangeName("Just Dano")
	nanodano.PrintInfo()
}
