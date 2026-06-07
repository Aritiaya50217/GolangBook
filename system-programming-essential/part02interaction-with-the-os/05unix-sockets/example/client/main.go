package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("unix", "/tmp/app.sock")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("ping"))

	buf := make([]byte, 1024)

	n, _ := conn.Read(buf)

	fmt.Println(string(buf[:n]))
}
