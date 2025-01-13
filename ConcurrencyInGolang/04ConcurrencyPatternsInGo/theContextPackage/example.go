package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // รอ (block) จนกว่า Context จะถูกยกเลิกหรือหมดเวลา
			fmt.Println("stopped : ", ctx.Err())
			return
		default:
			fmt.Println(name, "working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func withCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "Worker 1")
	go worker(ctx, "Worker 2")

	time.Sleep(2 * time.Second)

	fmt.Println("canceling context...")
	cancel() // ยกเลิก context

	time.Sleep(1 * time.Second) // รอให้ goroutines หยุด

}

func withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // สร้าง Context ใหม่ที่หมดเวลาอัตโนมัติเมื่อถึง Timeout ที่กำหนด
	defer cancel()

	go worker(ctx, "Worker 1")
	time.Sleep(3 * time.Second)
}

func detail(ctx context.Context) {
	userId := ctx.Value("userId")
	fmt.Println("User ID : ", userId)
}

func withValue() {
	ctx := context.WithValue(context.Background(), "userId", 42)
	detail(ctx)
}

func main() {
	withCancel()
}
