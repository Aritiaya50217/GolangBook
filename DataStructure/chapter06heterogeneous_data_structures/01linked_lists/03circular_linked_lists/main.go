package main

import "fmt"

// circular queue
type CircularQueue struct {
	size  int
	nodes []interface{}
	head  int
	last  int
}

func NewQueue(num int) *CircularQueue {
	circularQueue := CircularQueue{
		size: num + 1,
		head: 0,
		last: 0,
	}
	circularQueue.nodes = make([]interface{}, circularQueue.size)
	return &circularQueue
}

func (circularQueue CircularQueue) IsUnUsed() bool {
	return circularQueue.head == circularQueue.last
}

func (circularQueue CircularQueue) IsComplete() bool {
	return circularQueue.head == (circularQueue.last+1)%circularQueue.size
}

func (circularQueue *CircularQueue) Add(element interface{}) {
	if circularQueue.IsComplete() {
		panic("Queue is Completely Utilized")
	}
	circularQueue.nodes[circularQueue.last] = element
	circularQueue.last = (circularQueue.last + 1) % circularQueue.size
}

func (circularQueue *CircularQueue) MoveOneStep() (element interface{}) {
	if circularQueue.IsUnUsed() {
		return nil
	}

	element = circularQueue.nodes[circularQueue.head]
	circularQueue.head = (circularQueue.head + 1) % circularQueue.size
	return
}

func main() {
	var circularQueue *CircularQueue
	circularQueue = NewQueue(5)

	for i := 1; i <= 5; i++ {
		circularQueue.Add(i)
		fmt.Printf("Add(%d) : %v\n", i, circularQueue.nodes)
	}
}
