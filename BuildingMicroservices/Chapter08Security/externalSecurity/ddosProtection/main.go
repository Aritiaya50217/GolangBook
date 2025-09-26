package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var visitors = make(map[string]time.Time)
var mu sync.Mutex

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		mu.Lock()
		lastVisit, exists := visitors[ip]
		if exists && time.Since(lastVisit) < 1*time.Second {
			mu.Unlock()
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		visitors[ip] = time.Now()
		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", limitMiddleware(http.HandlerFunc(handler)))

	http.ListenAndServe(":8080", mux)
}
