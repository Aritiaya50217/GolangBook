package main

import "fmt"

// Originator
type Originator struct {
	state string
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	fmt.Println("Setting state to " + state)
	o.state = state
}

func (o *Originator) GetMemento() Memento {
	// externalize state to memento object
	return Memento{o.state}
}

func (o *Originator) Restore(memento Memento) {
	// restore state
	o.state = memento.GetState()
}

// Memento
type Memento struct {
	serializedState string
}

func (m *Memento) GetState() string {
	return m.serializedState
}

// caretaker
func Caretaker() {
	theOriginator := Originator{"A"}
	theOriginator.SetState("A")
	fmt.Println("theOriginator state = ", theOriginator.GetState())

	// before mutating  , get an memento
	theMemento := theOriginator.GetMemento()

	// mutate to unclean
	theOriginator.SetState("unclean")
	fmt.Println("theOriginator state = ", theOriginator.GetState())

	// rollback
	theOriginator.Restore(theMemento)
	fmt.Println("Restore : theOriginator state = ", theOriginator.GetState())

}

func main() {
	Caretaker()

}
