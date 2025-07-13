package main

import (
	"errors"
	"fmt"
)

// The Command - encapsulates the action to be done
type Report interface {
	Excute() // here the action is called Execute
}

// The Concrete Commands
type ConcreteReportA struct {
	// The action needs to be done on this receiver object
	receiver *Receiver
}

func (c *ConcreteReportA) Excute() {
	c.receiver.Action("ReportA")
}

type ConcreteReportB struct {
	receiver *Receiver
}

func (c *ConcreteReportB) Excute() {
	c.receiver.Action("ReportB")
}

// end of concrete commands

// Receiver - ancilly objects passed to command execution
// This can pass useful information for

type Receiver struct{}

func (r *Receiver) Action(msg string) {
	fmt.Println(msg)
}

// Invoker - this object which knows how to execute a command, and optionally
// does bookkeeping about the command execution.
type Invoker struct {
	repository []Report
}

func (i *Invoker) Schedule(cmd Report) {
	i.repository = append(i.repository, cmd)
}

func (i *Invoker) Run() {
	for _, cmd := range i.repository {
		cmd.Excute()
	}
}

// Chain of Responsibilty
// uses Command to represent request as objects
type ChainedReceiver struct {
	canHandle string
	next      *ChainedReceiver
}

func (r *ChainedReceiver) SetNext(next *ChainedReceiver) {
	r.next = next
}

func (r *ChainedReceiver) Finish() error {
	fmt.Println(r.canHandle, " Receiver Finishing")
	return nil
}

func (r *ChainedReceiver) Handle(what string) error {
	// Check if this receiver can handle
	// this of course is a dummy check
	if what == r.canHandle {
		return r.Finish()
	} else if r.next != nil {
		return r.next.Handle(what)
	} else {
		fmt.Println("No Receiver could handle the request!")
		return errors.New("No Receiver to Handle")
	}
}

func main() {
	receiver := new(Receiver)
	reportA := &ConcreteReportA{receiver: receiver}
	reportB := &ConcreteReportB{receiver: receiver}
	invoker := new(Invoker)
	invoker.Schedule(reportA)
	invoker.Run()
	invoker.Schedule(reportB)
	invoker.Run()

}
