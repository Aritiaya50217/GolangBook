package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		cancel() // ยกเลิก context หลัง 2 วินาที
	}()

	process(ctx)
}

func process(ctx context.Context) {
	fmt.Println("เริ่มทำงาน...")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("ทำงานเสร็จ")
	case <-ctx.Done():
		fmt.Println("ยกเลิกการทำงาน: ", ctx.Err())
	}
}
