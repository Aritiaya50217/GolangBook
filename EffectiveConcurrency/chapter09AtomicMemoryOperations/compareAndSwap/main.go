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
	myCopy := sheredValue.Load()       // โหลดค่าปัจจุบันของ shared struct แบบ atomic
	newCopy := computeNewCopy(*myCopy) // copy ใหม่ โดยเพิ่มค่า v
	
	// พยายามอัปเดต shared pointer
	if sheredValue.CompareAndSwap(myCopy, &newCopy) { // ถ้า sheredValue ยังชี้ไปที่ myCopy อยู่ → เปลี่ยนให้ชี้ไปที่ newCopy แล้วคืน true
		fmt.Printf("Set value %d\n", index)
	} else {
		fmt.Printf("Cannot set value %d\n", index)
	}
}

func main() {
	sheredValue.Store(&SomeStruct{}) // กำหนดค่าเริ่มต้นของ shared struct เป็น {v: 0}
	wg := sync.WaitGroup{}
	// สร้าง goroutine 2 ตัว (วนลูป 0 ถึง 1) ให้เรียก updateSheredValue(i)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			updateSheredValue(i)
		}()
	}
	wg.Wait()
}
