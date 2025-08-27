package main

import "fmt"

type Event struct {
	Message string
}

// Subscriber ที่รอรับ event
type Subscriber chan Event

// Publisher ที่ broadcast event
type Publisher struct {
	subscribers []Subscriber
}

func (p *Publisher) Subscribe() Subscriber {
	ch := make(Subscriber, 1)
	p.subscribers = append(p.subscribers, ch)
	return ch
}

func (p *Publisher) Publisher(event Event) {
	for _, sub := range p.subscribers {
		sub <- event
	}
}

func main() {
	p := &Publisher{}
	sub1 := p.Subscribe()
	sub2 := p.Subscribe()

	go func() {
		for e := range sub1 {
			fmt.Println("Sub1 received: ", e.Message)
		}
	}()

	go func() {
		for e := range sub2 {
			fmt.Println("Sub2 received: ", e.Message)
		}
	}()

	p.Publisher(Event{Message: "Hello Event!"})
	p.Publisher(Event{Message: "Another Event!"})
}
