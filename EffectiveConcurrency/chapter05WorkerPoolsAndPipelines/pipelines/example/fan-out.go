package main

import (
	"fmt"
	"math/rand"
	"time"
)

// producer : ส่งเลขเข้า channel
func generator() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// worker: ทำงานแบบ fan-out รับจาก input เดียว
func worker(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			fmt.Printf("Worker %d processing %d\n", id, n)
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	input := generator()

	// fan-out : สร้าง worker หลายตัวใช้ input เดียวกัน
	w1 := worker(1, input)
	w2 := worker(2, input)

	// อ่าน output จากแต่ละ worker
	for v := range merge(w1, w2) {
		fmt.Println("Result:", v)
	}
}

// merge : รวมผลจากหลาย channel (fan-in)
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
		}(c)
	}
	return out
}
