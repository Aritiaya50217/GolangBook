package main

import (
	"log"
	"net/http"
	"test/web"
)

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", web.Home)
	http.ListenAndServe(":8000", nil)
}
