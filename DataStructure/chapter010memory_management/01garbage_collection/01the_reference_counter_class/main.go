package main

import (
	"fmt"
	"sync"
)

type ReferenceCounter struct {
	num     *uint32
	pool    *sync.Pool
	removed *uint32
}

func newReferenceCounter() *ReferenceCounter {
	return &ReferenceCounter{
		num:     new(uint32),
		pool:    &sync.Pool{},
		removed: new(uint32),
	}
}

type Stack struct {
	references []*ReferenceCounter
	Count      int
}

func (stack *Stack) New() {
	stack.references = make([]*ReferenceCounter, 0)
}

func (stack *Stack) Push(reference *ReferenceCounter) {
	stack.references = append(stack.references[:stack.Count], reference)
	stack.Count += 1
	fmt.Println("Pop : ", stack.references)
}

func (stack *Stack) Pop() *ReferenceCounter {
	if stack.Count == 0 {
		return nil
	}

	length := len(stack.references)
	reference := stack.references[length-1]

	if length > 1 {
		stack.references = stack.references[:length-1]
	} else {
		stack.references = stack.references[0:]
	}

	stack.Count = len(stack.references)
	return reference
}

func main() {
	stack := &Stack{}
	stack.New()

	reference1 := newReferenceCounter()
	reference2 := newReferenceCounter()
	reference3 := newReferenceCounter()
	reference4 := newReferenceCounter()

	stack.Push(reference1)
	stack.Push(reference2)
	stack.Push(reference3)
	stack.Push(reference4)

	fmt.Println("==== Pop ====")
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())

}
