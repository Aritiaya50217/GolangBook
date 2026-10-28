package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ReferenceCounter struct {
	num     *uint32
	pool    *sync.Pool
	removed *uint32
}

func newReferenceCounter() ReferenceCounter {
	return ReferenceCounter{
		num:     new(uint32),
		pool:    &sync.Pool{},
		removed: new(uint32),
	}
}

func (referenceCounter ReferenceCounter) Add() {
	atomic.AddUint32(referenceCounter.num, 1)
}

func (referenceCounter ReferenceCounter) Subtract() {
	if atomic.AddUint32(referenceCounter.num, ^uint32(0)) == 0 {
		atomic.AddUint32(referenceCounter.removed, 1)
	}
}

func main() {
	referenceCounter := newReferenceCounter()
	referenceCounter.Add()
	referenceCounter.Add()
	referenceCounter.Add()
	fmt.Println("before : ", *referenceCounter.num)

	referenceCounter.Subtract()
	fmt.Println("after : ", *referenceCounter.num)
}
