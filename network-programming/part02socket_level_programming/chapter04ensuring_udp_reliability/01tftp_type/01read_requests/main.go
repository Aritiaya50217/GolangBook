package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:69")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	filename := "test.txt"
	mode := "octet"

	// RRQ packet
	packet := append(
		[]byte{0, 1}, // Opcode = RRQ
		[]byte(filename)...,
	)

	packet = append(packet, 0)
	packet = append(packet, []byte(mode)...)
	packet = append(packet, 0)

	if _, err := conn.Write(packet); err != nil {
		panic(err)
	}

	fmt.Println("RRQ sent")
}
