package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// ส่งข้อมูลไปหา echo server
	message := []byte("Hello UDP")

	_, err = conn.Write(message)
	if err != nil {
		panic(err)
	}

	// เตรียม buffer สำหรับรับข้อมูล
	buffer := make([]byte, 1024)

	// รับข้อมูลที่ Echho Server ส่งกลับมา
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("Received:", string(buffer[:n]))
}
