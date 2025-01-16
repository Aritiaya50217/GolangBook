package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second): // งานใช้เวลา 3 วินาที
		return fmt.Errorf("task complated")
	case <-ctx.Done(): // ถูกยกเลิกหรือ Timeout
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // ยกเลิก Context เมื่อฟังก์ชันจบ

	if err := longRunningTask(ctx); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Task completed successfully")
	}

}
