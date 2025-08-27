package main

import (
	"fmt"
)

type Event interface{}

type UserCreated struct {
	ID   string
	Name string
}

type UserRenamed struct {
	ID   string
	Name string
}

type User struct {
	ID   string
	Name string
}

// command handler
func CreateUser(id, name string, bus chan<- Event) {
	bus <- UserCreated{ID: id, Name: name}
}

// Event handler
func HandlerUserCreated(e UserCreated) {
	fmt.Println("Updating read model: ", e)
}

func main() {
	bus := make(chan Event)

	go func() {
		for e := range bus {
			switch ev := e.(type) {
			case UserCreated:
				HandlerUserCreated(ev)
			}
		}
		fmt.Println("bus : ", bus)
	}()

	CreateUser("1", "Alice", bus)

	close(bus)
}
