package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":8080")
	conn, _ := net.ListenUDP("udp", addr)

	fmt.Println("Server:", conn.LocalAddr())

	buf := make([]byte, 1024)

	for {
		n, clientAddr, _ := conn.ReadFromUDP(buf)
		fmt.Printf(
			"Received: %s from %v\n",
			string(buf[:n]),
			clientAddr,
		)
	}
}
