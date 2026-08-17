package main

import (
	"fmt"
	"net/http"
	"time"
)

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "Done")
}

func main() {
	timeHandler := http.TimeoutHandler(http.HandlerFunc(slowHandler), 2*time.Second, "Request timed out\n")
	http.ListenAndServe(":8080", timeHandler)
}
