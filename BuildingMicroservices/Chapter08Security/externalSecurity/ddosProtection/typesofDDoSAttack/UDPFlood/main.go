package main

import (
	"log"
	"net"
)

func main() {
	target := "127.0.0.1:8080"
	conn, _ := net.Dial("udp", target)
	defer conn.Close()

	payload := []byte("Hello UDP Flood")
	for i := 0; i < 10000; i++ {
		conn.Write(payload)
	}
	log.Println("Finished sending UDP packets")
}
