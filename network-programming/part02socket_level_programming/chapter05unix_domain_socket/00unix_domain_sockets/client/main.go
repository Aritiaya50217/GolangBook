package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("Hello Server"))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	fmt.Println(string(buf[:n]))
}
