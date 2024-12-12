package main

import (
	"container/heap"
	"fmt"
)

// integerHeap a type
type IntegerHeap []int

// InegerHeap method - gets the length of integerHeap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

// IntegerHeap method - checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

// IntegerHeap method -swaps the element of i to j index
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

// IntegerHeap method -pushes the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

// IntegerHeap method -pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var previous IntegerHeap = *iheap
	n := len(previous)
	x1 := previous[n-1]
	fmt.Println("x1 : ", x1)
	*iheap = previous[0 : n-1]
	fmt.Println("*iheap : ", *iheap)
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	fmt.Println("Be : ", intHeap)
	heap.Init(intHeap)
	fmt.Println("Init : ", intHeap)
	heap.Push(intHeap, 2)
	fmt.Println("push : 2 ", intHeap)
	fmt.Println("minimum : ", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("value : %v \n--------------------\n", heap.Pop(intHeap))
	}
}
