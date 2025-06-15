package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func makeItSpeak(s Speaker) {
	s.Speak()
}
func main() {

	makeItSpeak(Dog{})
	makeItSpeak(Cat{})

}
