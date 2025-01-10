package main

import "fmt"

func producer(id int, out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- id*10 + i
	}
}

func fanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	for _, in := range inputs {
		go func(ch <-chan int) {
			defer close(out) // ปิด channels หลังส่งข้อมูลเสร็จ
			for val := range ch {
				out <- val
			}
		}(in)
	}

	return out
}

func main() {

	// สร้าง Producers
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go producer(1, ch1)
	go producer(2, ch2)
	go producer(3, ch3)

	// รวมผลลัพธ์ด้วย fan in
	result := fanIn(ch1, ch2, ch3)

	for i := 0; i < 9; i++ {
		fmt.Println("Result : ", <-result)
	}
}
