package main

import (
	"fmt"
	"time"
)

type Semaphore chan struct{}

// สร้าง semaphore แบบไม่มี buffer เริ่มต้น (counter = 0)
func NewSemaphore() Semaphore {
	return make(Semaphore, 1)
}

// ส่งสัญญาณ (increment)
func (s Semaphore) Signal() {
	select {
	case s <- struct{}{}:
	default:
		// ถ้าเต็มแล้วก็ไม่ต้องส่งเพิ่ม
	}
}

// รอรับสัญญาณ (decrement)
func (s Semaphore) Wait() {
	<-s
}

func main() {
	s := NewSemaphore()

	// worker มาช้า 2 วิ
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Worker waiting ...")
		s.Wait() // จะไม่พลาดสัญญาณ
		fmt.Println("Worker received signal ..")
	}()

	// main ส่งสัญญาณก่อน worker จะพร้อม
	fmt.Println("Main sending signal first")
	s.Signal() // signal ถูกส่งก่อนแต่ไม่หาย
	fmt.Println("Main done")

	time.Sleep(3 * time.Second)
}
