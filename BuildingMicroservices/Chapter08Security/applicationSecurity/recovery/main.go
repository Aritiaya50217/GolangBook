package main

import (
	"log"
	"net/http"
)

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[ALERT] Panic recovered: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	panic("Boom!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	port := "8080"

	log.Println("Server running at : " + port)
	http.ListenAndServe(port, recoveryMiddleware(mux))

}
