package main

import (
	"context"
	"fmt"
	"time"
)

func cancellableTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled: ", ctx.Err())
			return
		default:
			fmt.Println("Task running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // Context ที่สามารถยกเลิกได้

	go cancellableTask(ctx)
	time.Sleep(2 * time.Second) // ให้ Task ทำงาน 2 วินาที
	cancel()                    // ส่งสัญญาณยกเลิก
	time.Sleep(1 * time.Second) // รอ Task หยุดทำงาน

}
