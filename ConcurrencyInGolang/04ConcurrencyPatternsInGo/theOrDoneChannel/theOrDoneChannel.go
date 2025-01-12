package main

import (
	"fmt"
	"time"
)

// orDone ฟังก์ชันสร้าง Channel ที่จะหยุดการทำงานเมื่อ "done" ถูกปิด
func orDone(done <-chan struct{}, input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done: // หยุดทำงานเมื่อ done ถูกปิด
				return
			case val, ok := <-input: // อ่านค่า input channel
				if !ok {
					return // หยุดเมื่อ input ถูกปิด
				}
				out <- val // ส่งค่าต่อไปยัง output channel
			}
		}
	}()
	return out
}

func singleChannel() {
	done := make(chan struct{})
	input := make(chan int)

	// Goroutine ที่ส่งค่าผ่าน input channel
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(input)
	}()

	// goroutine สำหรับยกเลิก (close done channel)
	go func() {
		time.Sleep(250 * time.Millisecond) // ยกเลิกหลัง 250 ms
		close(done)
	}()

	// อ่านค่าจาก orDone channel
	for val := range orDone(done, input) {
		fmt.Println("input : ", val)
	}
	fmt.Println("Processing stopped")
}

func or(channels ...<-chan struct{}) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		defer close(out)
		for _, ch := range channels {
			go func(c <-chan struct{}) {
				<-c
				out <- struct{}{}
			}(ch)
		}
	}()
	return out
}
func multipleChannel() {
	done1 := make(chan struct{})
	done2 := make(chan struct{})
	combined := or(done1, done2)

	go func() {
		close(done1) // ส่งสัญญาณยกเลิก
	}()

	<-combined // รอจนกว่าสัญญาณจากช่องใดช่องหนึ่งจะมาถึง

}

func main() {
	// singleChannel()
	multipleChannel()
}
