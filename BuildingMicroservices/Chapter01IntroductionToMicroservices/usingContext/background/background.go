package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("ทำงานเสร็จ")
	case <-ctx.Done():
		fmt.Println("หมดเวลา : ", ctx.Err())
	}

}
