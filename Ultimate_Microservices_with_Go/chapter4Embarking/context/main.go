package main

import (
	"context"
	"fmt"
	"time"
)

func printAPIKey(ctx context.Context) {
	apiKey := ctx.Value("apiKey")
	fmt.Println("API Key: ", apiKey)
}

func main() {
	ctx := context.Background()
	childContext := context.WithValue(ctx, "apiKey", 123456)
	printAPIKey(childContext)

	contextWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	fmt.Println("What happend ?", contextWithCancel.Err())

	contextWithTimeout, _ := context.WithTimeout(ctx, 30*time.Second)
	fmt.Println("What happened ?", contextWithTimeout.Err())
	time.Sleep(35 * time.Second)
	fmt.Println("What happened ?", contextWithTimeout.Err())

}
