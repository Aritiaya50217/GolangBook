package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "สวัสดี นี่คือการเชื่อมต่อแบบ TLS!")
}

func main() {
	http.HandleFunc("/", handler)

	// ต้องมีไฟล์ certificate (.cert) และ priveate key (.key)
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
