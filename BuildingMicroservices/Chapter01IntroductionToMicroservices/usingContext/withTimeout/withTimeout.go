package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	process(ctx)
}

func process(ctx context.Context) {
	fmt.Println("start ...")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("end")
	case <-ctx.Done():
		fmt.Println("cancel: ", ctx.Err())
	}
}
