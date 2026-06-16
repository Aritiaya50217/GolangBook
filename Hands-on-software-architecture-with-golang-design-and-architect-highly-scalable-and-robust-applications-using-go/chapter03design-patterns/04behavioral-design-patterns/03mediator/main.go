package main

import (
	"fmt"
	"strconv"
)

// The Mediator interface
type Mediator interface {
	AddColleague(colleague Colleague)
}

// The Colleague interface
type Colleague interface {
	setMediator(mediator Mediator)
}

// Concrete Colleague 1 - uses state as string
type Colleague1 struct {
	mediator Mediator
	state    string
}

func (c *Colleague1) setMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *Colleague1) SetState(state string) {
	fmt.Println("Colleague1: setting state: ", state)
	c.state = state
}

func (c *Colleague1) GetState() string {
	return c.state
}

type Colleague2 struct {
	mediator Mediator
	state    int
}

func (c *Colleague2) SetState(state int) {
	fmt.Println("Colleague2: setting state: ", state)
	c.state = state
}

func (c *Colleague2) GetState() int {
	return c.state
}

func (c *Colleague2) setMediator(mediator Mediator) {
	c.mediator = mediator
}

// Concreate Mediator
type ConcreteMediator struct {
	c1 Colleague1
	c2 Colleague2
}

func (m *ConcreteMediator) SetColleagueC1(c1 Colleague1) {
	m.c1 = c1
}
func (m *ConcreteMediator) SetColleagueC2(c2 Colleague2) {
	m.c2 = c2
}

func (m *ConcreteMediator) SetState(s string) {
	m.c1.SetState(s)
	stateAsString, err := strconv.Atoi(s)
	if err == nil {
		m.c2.SetState(stateAsString)
		fmt.Println("Mediator set status for both colleagues")
	}
}
func main() {
	c1 := Colleague1{}
	c2 := Colleague2{}

	// initialize mediator with colleagues
	mediator := ConcreteMediator{}
	mediator.SetColleagueC1(c1)
	mediator.SetColleagueC2(c2)

	// mediator keeps colleagues in sync
	mediator.SetState("10")

}
