package main

import (
	"context"
	"fmt"
	"time"
)

func mixedTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // งานใช้เวลา 5 วินาที
		fmt.Println("Task completed")
	case <-ctx.Done(): // Timeout หรือ Cancelled
		fmt.Println("Task cancelled : ", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //  // Timeout 3 วินาที
	defer cancel()

	go mixedTask(ctx)

	time.Sleep(2 * time.Second) // รอ 2 วินาที
	cancel()   // ยกเลิก Context ก่อน Timeout
	time.Sleep(1 * time.Second) // รอให้ Task หยุด

}
