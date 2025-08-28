package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

func main() {
	settings := gobreaker.Settings{
		Name:        "HTTP GET",
		MaxRequests: 3,                // number of requests allowed in half-open state
		Interval:    60 * time.Second, // interval to reset counts
		Timeout:     10 * time.Second, // how long to stay before trying again
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// trip circuit if > 5 failures
			return counts.ConsecutiveFailures > 5
		},
	}
	cb := gobreaker.NewCircuitBreaker(settings)

	for i := 0; i < 10; i++ {
		_, err := cb.Execute(func() (interface{}, error) {
			resp, err := http.Get("http://localhost:8080/health")
			// resp, err := http.Get("https://www.google.com") // ตัวอย่าง success
			if err != nil || resp.StatusCode >= 500 {
				return nil, fmt.Errorf("service unavailable")
			}
			return resp, nil
		})
		if err != nil {
			fmt.Println("Request failed : ", err)
		} else {
			fmt.Println("Request success")
		}
		time.Sleep(1 * time.Second)
	}

}
