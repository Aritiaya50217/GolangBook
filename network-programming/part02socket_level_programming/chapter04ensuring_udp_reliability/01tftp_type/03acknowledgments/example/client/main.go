package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	OpRRQ  = 1
	OpDATA = 3
	OpACK  = 4
)

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:9000")
	conn, _ := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()

	// ส่ง RRQ แบบง่าย
	conn.Write([]byte("test.txt"))

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	opcode := binary.BigEndian.Uint16(buf[0:2])
	block := binary.BigEndian.Uint16(buf[2:4])
	data := buf[4:n]

	if opcode == OpDATA {
		fmt.Printf("Received DATA block #%d: %s\n", block, string(data))

		// สร้าง ACK Packet
		ack := make([]byte, 4)
		binary.BigEndian.PutUint16(ack[0:2], OpACK)
		binary.BigEndian.PutUint16(ack[2:4], block)

		conn.Write(ack)
		fmt.Printf("Sent ACK block #%d\n", block)
	}
}
