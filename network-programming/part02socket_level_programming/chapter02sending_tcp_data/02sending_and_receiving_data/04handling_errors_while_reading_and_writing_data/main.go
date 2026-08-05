package main

import (
	"errors"
	"fmt"
	"io"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
				return
			}

			var netErr net.Error

			if errors.As(err, &netErr) && netErr.Timeout() {
				fmt.Println("Read timeout")
				return
			}

			fmt.Println("Read error : ", err)
			return
		}

		if _, err := conn.Write(buf[:n]); err != nil {
			fmt.Println("Write error : ", err)
			return
		}
	}
}

func main() {

}
