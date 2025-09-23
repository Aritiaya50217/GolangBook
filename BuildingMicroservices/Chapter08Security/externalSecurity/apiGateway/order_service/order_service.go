package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Order service response: %s", r.URL.Path)
	})

	log.Println("Order service running on :9002")
	log.Fatal(http.ListenAndServe(":9002", nil))
}
