package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// สร้าง request จำลอง
	req := httptest.NewRequest("GET", "/hello?name=Golang", nil)

	// สร้าง response recorder (จำลอง response writer)
	w := httptest.NewRecorder()

	// เรียก handler
	HelloHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}
