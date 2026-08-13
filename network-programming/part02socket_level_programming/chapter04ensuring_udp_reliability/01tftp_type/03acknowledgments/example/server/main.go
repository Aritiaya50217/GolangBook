package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	OpDATA = 3
	OpACK  = 4
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":9000")
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}

	defer conn.Close()

	fmt.Println("TFTP-like Server listening on :9000")

	buf := make([]byte, 1024)

	// รอ RRQ
	n, clientAddr, _ := conn.ReadFromUDP(buf)
	fmt.Printf("Received request : %q\n", string(buf[:n]))

	// สร้าง DATA Packet (Block #1)
	data := []byte("Hello TFTP")

	packet := make([]byte, 4+len(data))
	binary.BigEndian.PutUint16(packet[0:2], OpDATA)
	binary.BigEndian.PutUint16(packet[2:4], 1)
	copy(packet[4:], data)

	conn.WriteToUDP(packet, clientAddr)
	fmt.Println("Sent DATA block #1")

	// รอ ACK
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	n, _, err = conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("Timeout waiting for ACK")
		return
	}

	opcode := binary.BigEndian.Uint16(buf[0:2])
	block := binary.BigEndian.Uint16(buf[2:4])

	if opcode == OpACK && block == 1 {
		fmt.Println("Received ACK block #1")
	}
}
