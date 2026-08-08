package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":9000")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("UDP Server listening on :9000")

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received from %s: %s\n", clientAddr, string(buffer[:n]))

		// ส่งข้อมูลกลับไปหา Client
		if _, err := conn.WriteToUDP([]byte("Hello from UDP Server"), clientAddr); err != nil {
			panic(err)
		}

	}
}
