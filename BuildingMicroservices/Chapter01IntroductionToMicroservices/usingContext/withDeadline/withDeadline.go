package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// กำหนด deadline ว่าจะหมดอายุในอีก 3 วินาที
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	process(ctx)
}

func process(ctx context.Context) {
	fmt.Println("start ...")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("end")
	case <-ctx.Done():
		fmt.Println("cancel : ", ctx.Err())
	}
}
