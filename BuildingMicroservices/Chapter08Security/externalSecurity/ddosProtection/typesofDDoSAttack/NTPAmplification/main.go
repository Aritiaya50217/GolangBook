package main

import (
	"log"
	"net"
)

func main() {
	target := "127.0.0.1:123" // NTP port
	conn, _ := net.Dial("udp", target)
	defer conn.Close()

	payload := []byte{0x17, 0x00, 0x03, 0x2a} // monlist request format
	for i := 0; i < 100; i++ {
		conn.Write(payload)
		log.Println("Sent NTP monlist request")
	}
}
