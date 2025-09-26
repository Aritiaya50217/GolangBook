package main

import (
	"log"
	"net"
)

func main() {
	target := "127.0.0.1:19" // Chargen port
	conn, _ := net.Dial("udp", target)
	defer conn.Close()

	payload := []byte("test")
	for i := 0; i < 1000; i++ {
		conn.Write(payload)
		log.Println("Sent Chargen request")
	}
}
