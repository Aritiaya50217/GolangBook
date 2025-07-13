package main

import (
	"errors"
	"fmt"
)

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
	if what == r.canHandle {
		return r.Finish()
	} else if r.next != nil {
		// delegate to the next guy
		return r.next.Handle(what)
	} else {
		fmt.Println("No Receiver could handle the request!")
		return errors.New("No Receiver to Handl")
	}
}

func main() {
	var chainedReceiver ChainedReceiver
	data := ChainedReceiver{
		canHandle: "Text",
	}

	chainedReceiver.SetNext(&data)
	chainedReceiver.Handle("Text")
}
