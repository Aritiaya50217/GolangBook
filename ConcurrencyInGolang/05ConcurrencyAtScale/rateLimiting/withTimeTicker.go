package main

import (
	"fmt"
	"time"
)

func main() {
	rate := time.Second / 2 // จำกัด 2 คำขอต่อวินาที
	ticker := time.NewTicker(rate)

	defer ticker.Stop()

	requests := []string{"req1", "req2", "req3", "req4", "req5"}

	for _, req := range requests {
		<-ticker.C // รอให้ Ticker ส่งสัญญาณ
		fmt.Println("Processing", req, "at", time.Now())
	}

}
