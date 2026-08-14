package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

const OpERROR = 5 // Opcode = 5 หมายถึง ERROR

func main() {
	conn, _ := net.Dial("udp", "localhost:9000")
	defer conn.Close()

	conn.Write([]byte("missing.txt"))

	buf := make([]byte, 1024)

	n, _ := conn.Read(buf)

	opcode := binary.BigEndian.Uint16(buf[0:2])

	if opcode == OpERROR {
		code := binary.BigEndian.Uint16(buf[2:4])
		message := string(buf[4 : n-1])

		fmt.Printf("TFTP Error %d: %s\n", code, message)
	}
}
