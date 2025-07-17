package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// handler ที่จะทดสอบ
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello , test!"))
}

func main() {
	// สร้าง mock request
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	// สร้าง mocl reqponse writer
	rec := httptest.NewRecorder()

	// เรียก handler โดยใช้ mock request และ response
	HelloHandler(rec, req)

	// ดึง response ที่เขียนโดย handler ออกมา
	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println("Status Code:", res.StatusCode)                  // Status Code: 200
	fmt.Println("Content-Type:", res.Header.Get("Content-Type")) // Content-Type: text/plain
	fmt.Println("Body:", string(body))

}
