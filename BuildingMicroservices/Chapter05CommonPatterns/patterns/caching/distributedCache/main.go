package main

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// set cache with TTL 10s
	ctx := context.Background()
	err := rdb.Set(ctx, "user:1", "Alice", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "user:1").Result()
	if err == redis.Nil {
		fmt.Println("cache miss")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("From redis cache : ", val)
	}
}
