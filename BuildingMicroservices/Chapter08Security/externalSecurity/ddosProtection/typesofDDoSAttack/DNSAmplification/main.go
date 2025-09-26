package main

import (
	"log"
	"net"
	"time"
)

func main() {
	server := "8.8.8.8:53"
	msg := []byte{}

	conn, _ := net.Dial("udp", server)
	defer conn.Close()

	for i := 0; i < 100; i++ {
		conn.Write(msg)
		time.Sleep(10 * time.Millisecond)
		log.Println("Sent DNS query")
	}
}
