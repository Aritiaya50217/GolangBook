package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SomeStruct struct {
	v int
}

var sheredValue atomic.Pointer[SomeStruct]

func computeNewCopy(in SomeStruct) SomeStruct {
	return SomeStruct{v: in.v + 1}
}

func updateSheredValue(index int) {
	myCopy := sheredValue.Load()
	newCopy := computeNewCopy(*myCopy)
	if sheredValue.CompareAndSwap(myCopy, &newCopy) {
		fmt.Printf("Set value %d\n", index)
	} else {
		fmt.Printf("Cannot set value %d\n", index)
	}
}

func main() {
	sheredValue.Store(&SomeStruct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			updateSheredValue(i)
		}()
	}
	wg.Wait()
}
