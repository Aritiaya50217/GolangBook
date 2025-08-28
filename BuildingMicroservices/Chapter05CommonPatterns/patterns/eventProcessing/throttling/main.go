package main

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

func main() {
	// create limiter ที่อนุญาต 5 req/sec และ brust ได้ 10
	limiter := rate.NewLimiter(5, 10)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests : ", http.StatusTooManyRequests)
			return
		}
		fmt.Fprintln(w, "Hello , world")
	})
	fmt.Println("Server started at : 8080")
	http.ListenAndServe(":8080", nil)
}
