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

	buffer := make([]byte, 1024)

	for {
		// รับข้อมูลจาก Client
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		// แสดงข้อมูลที่ได้รับ
		fmt.Println("Received:", string(buffer[:n]))
		fmt.Println("From:", clientAddr)

		// ส่งข้อมูลเดิมกลับไป
		_, err = conn.WriteToUDP(buffer[:n], clientAddr)
		if err != nil {
			panic(err)
		}
	}
}
