package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User service response : %s", r.URL.Path)
	})
	log.Println("User service running on : 9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
