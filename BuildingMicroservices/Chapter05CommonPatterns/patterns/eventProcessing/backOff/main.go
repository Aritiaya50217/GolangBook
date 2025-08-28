package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

func retryRequest(url string, maxRetries int, baseDelay time.Duration) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; i < maxRetries; i++ {
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode != http.StatusOK {
			return resp, nil
		}

		// exponential backoff
		backoff := time.Duration(math.Pow(2, float64(i))) * baseDelay
		fmt.Printf("Retry %d: waiting %v before retrying...\n", i+1, backoff)
		time.Sleep(backoff)

	}
	return resp, fmt.Errorf("failed after %d retries: %w", maxRetries, err)
}

func main() {
	_, err := retryRequest("https://google.com/", 5, 100*time.Millisecond)
	if err != nil {
		fmt.Println("Request failed:", err)
	}
}
