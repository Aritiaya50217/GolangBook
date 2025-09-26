package main

import (
	"log"
	"net"
)

func main() {
	target := "239.255.255.250:1900" // SSDP multicast
	conn, _ := net.Dial("udp", target)
	defer conn.Close()

	payload := []byte("M-SEARCH * HTTP/1.1\r\nHOST:239.255.255.250:1900\r\nMAN:\"ssdp:discover\"\r\nMX:1\r\nST:ssdp:all\r\n\r\n")
	for i := 0; i < 100; i++ {
		conn.Write(payload)
		log.Println("Sent SSDP discovery")
	}
}
