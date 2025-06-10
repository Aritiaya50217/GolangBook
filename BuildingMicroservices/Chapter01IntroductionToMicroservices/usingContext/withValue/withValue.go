package main

import (
	"context"
	"fmt"
)

type contextKey string

func main() {
	const userIDKey contextKey = "userID"

	// สร้าง context เปล่า
	ctx := context.Background()

	// ใส่ค่า userID เข้าไปใน context
	ctx = context.WithValue(ctx, userIDKey, 12345)
	// ดึงค่าออกมา
	userID := ctx.Value(userIDKey)
	fmt.Println("userID : ", userID)
}
