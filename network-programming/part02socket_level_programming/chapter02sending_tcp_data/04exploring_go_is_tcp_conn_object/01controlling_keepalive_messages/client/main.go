package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:9000")

	conn, err := net.DialTCP("tcp", nil, addr)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// เปิด Keepalive ฝั่ง Client
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(30 * time.Second)

	for {
		conn.Write([]byte("PING"))

		buf := make([]byte, 1024)

		n, _ := conn.Read(buf)

		fmt.Println(string(buf[:n]))

		time.Sleep(10 * time.Second)

	}
}
