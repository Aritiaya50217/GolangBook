package main

import "fmt"

// stage 1 : ส่งเลขจำนวนเต็มเข้าไป
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// stage 2 : คูณเลขด้วย 2
func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// stage 3 : แสดงผล
func main() {
	c := gen(1, 2, 3, 4)
	out := multiply(c)
	for n := range out {
		fmt.Println("res : ", n)
	}
}
