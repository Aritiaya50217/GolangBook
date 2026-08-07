package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9000")

	ln, _ := net.ListenTCP("tcp", addr)
	defer ln.Close()

	conn, _ := ln.AcceptTCP()

	if err := conn.SetReadBuffer(1024 * 1024); err != nil {
		panic(err)
	}

	fmt.Println("Read Buffer = 1MB")
}
