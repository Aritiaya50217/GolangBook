package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// middleware type: takes and returns http.Handler
type Middleware func(http.Handler) http.Handler

// Error handling middleware
func errorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Authentication middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")
		if token != "secret" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Final handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if strings.TrimSpace(name) == "" {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// pipeline builder
func chainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

func main() {
	finalHandler := http.HandlerFunc(helloHandler)
	pipeline := chainMiddleware(
		finalHandler,
		errorHandlerMiddleware,
		authMiddleware,
		loggingMiddleware,
	)
	http.Handle("/", pipeline)

	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
