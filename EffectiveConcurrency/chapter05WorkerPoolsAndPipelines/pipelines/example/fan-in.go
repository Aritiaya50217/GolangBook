package main

import (
	"fmt"
	"time"
)

// ทำงานเป็น producer แต่ละตัวจะส่งข้อมูลของตัวเองมา
func producer(id int, delay time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			out <- fmt.Sprintf("Producer %d: %d", id, i)
			time.Sleep(delay)
		}
		close(out)
	}()
	return out
}

// fan-in function : รวมหลาย channel เข้าเป็น channel เดียว
func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)
	for _, ch := range channels {
		go func(c <-chan string) {
			for v := range c {
				out <- v
			}
		}(ch)
	}
	return out
}

func main() {
	p1 := producer(1, time.Millisecond*500)
	p2 := producer(2, time.Millisecond*500)

	merged := fanIn(p1, p2)
	// รับค่าจาก channel เดียวที่รวมมาจากหลาย producer
	for i := 0; i < 6; i++ {
		fmt.Println(<-merged)
	}
}
