package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := ln.Accept()

		go func(c net.Conn) {
			defer c.Close()

			buf := make([]byte, 1024)
			n, _ := c.Read(buf)

			fmt.Println(string(buf[:n]))
		}(conn)
	}

}
