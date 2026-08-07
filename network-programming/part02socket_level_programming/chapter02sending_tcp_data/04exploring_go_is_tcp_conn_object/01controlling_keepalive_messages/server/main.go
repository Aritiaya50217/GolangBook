package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func handle(conn *net.TCPConn) {
	defer conn.Close()

	// เปิด Keepalive
	if err := conn.SetKeepAlive(true); err != nil {
		fmt.Println(err)
		return
	}

	// ส่ง Keepalive ทุก 30 s เมื่อไม่มีข้อมูล
	if err := conn.SetKeepAlivePeriod(30 * time.Second); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Client Connected: ", conn.RemoteAddr())

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client Closed")
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Println("Receive : ", string(buf[:n]))
		conn.Write([]byte("PONG"))
	}
}

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", ":9000")

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("Server Start")

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		go handle(conn)
	}
}
