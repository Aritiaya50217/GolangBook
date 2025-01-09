package main

import (
	"context"
	"fmt"
	"time"
)

// ขั้นตอนที่ 1: สร้างเลขจำนวนเต็ม
func generate(ctx context.Context, numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numbers {
			select {
			case <-ctx.Done(): // รอ (block) จนกว่า Context จะถูกยกเลิกหรือหมดเวลา
				return
			case out <- n:
			}
		}
	}()

	return out
}

// ขั้นตอนที่ 2: คูณด้วย 2
func multipleByTwo(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

// ขั้นตอนที่ 3: กรองตัวเลขที่มากกว่า 10
func filterGreaterThanTen(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n > 10 {
				out <- n
			}
		}
	}()
	return out
}

func pipelineWithContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	numbers := generate(ctx, 1, 2, 3, 4, 5, 6)

	// ส่งตัวเลขผ่าน Pipline
	multiplied := multipleByTwo(numbers)
	filtered := filterGreaterThanTen(multiplied)

	// แสดงผล
	for n := range filtered {
		fmt.Println(n)
	}
}

func main() {
	pipelineWithContext()
}
