package main

import (
	"container/list"
	"fmt"
)

func main() {
	linkedList := list.New()
	element := linkedList.PushBack(14)
	frontElement := linkedList.PushFront(1)
	linkedList.InsertBefore(6, element)
	linkedList.InsertAfter(5, frontElement)

	for currElement := linkedList.Front(); currElement != nil; currElement = currElement.Next() {
		fmt.Println(currElement.Value)
	}
}
