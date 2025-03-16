package main

import (
	"container/list"
	"fmt"
)

func main() {
	var inList list.List
	inList.PushBack(11)
	inList.PushBack(23)
	inList.PushBack(34)
	for element := inList.Front(); element != nil; element = element.Next() {
		fmt.Println("element : ", element.Value)
		fmt.Println("next : ", element.Next())
		fmt.Println("prev : ", element.Prev())
	}
}
