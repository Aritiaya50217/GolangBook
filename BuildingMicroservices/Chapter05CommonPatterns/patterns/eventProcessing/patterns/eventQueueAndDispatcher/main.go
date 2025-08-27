package main

import (
	"fmt"
)

type Event struct {
	Type string
	Data interface{}
}

func dispatcher(events <-chan Event) {
	for e := range events {
		switch e.Type {
		case "UserCreated":
			fmt.Println("Handle user created: ", e.Data)
		case "OrderPlaced":
			fmt.Println("Handle order placed: ", e.Data)
		}
	}
}

func main() {
	events := make(chan Event, 5)
	go dispatcher(events)

	events <- Event{Type: "UserCreated", Data: "Alice"}
	events <- Event{Type: "OrderPlaced", Data: "Order#123"}

	close(events)
}
