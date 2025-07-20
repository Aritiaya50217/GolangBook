package main

import (
	"container/ring"
	"fmt"
)

func main() {
	integers := []int{1, 3, 5, 7}
	circularList := ring.New(len(integers))

	for i := 0; i < circularList.Len(); i++ {
		circularList.Value = integers[i]
		circularList = circularList.Next()
	}

	circularList.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()

	for i := 0; i < circularList.Len(); i++ {
		fmt.Print(circularList.Value, ",")
		circularList = circularList.Prev()
	}

	fmt.Println()

	circularList = circularList.Move(2)
	circularList.Do(func(element interface{}) {
		fmt.Println(element, ",")
	})
	fmt.Println()
}
