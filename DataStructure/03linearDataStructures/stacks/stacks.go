package main

import (
	"fmt"
	"strconv"
)

// Element class
type Element struct {
	elementValue int
}

// String method on Element class
func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	elements     []*Element
	elementCount int
}

// NewStack returns a new stack
func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

// Push adds a node to the stack.
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	// fmt.Println("stack.elements[:stack.elementCount] : ", stack.elements[:stack.elementCount])
	stack.elementCount++
}

// Pop removes and returns a node from the stack in last to first order.
func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}
	length := len(stack.elements)
	element := stack.elements[len(stack.elements)-1] // last element

	if length > 1 {
		stack.elements = stack.elements[:length-1]
		fmt.Println("stack.elements[:length-1] : ", stack.elements)
	} else {
		stack.elements = stack.elements[0:]
		fmt.Println("stack.elements[0:] : ", stack.elements[0:])
	}
	stack.elementCount = len(stack.elements)
	return element
}

func main() {
	stack := Stack{}
	stack.New()
	element1 := Element{
		elementValue: 3,
	}
	element2 := Element{
		elementValue: 5,
	}
	element3 := Element{
		elementValue: 7,
	}
	element4 := Element{
		elementValue: 9,
	}

	stack.Push(&element1)
	stack.Push(&element2)
	stack.Push(&element3)
	stack.Push(&element4)

	// pop
	stack.Pop() // pop 9
	stack.Pop() // pop 7

	fmt.Println("element in stack : ", stack.elements)

}
