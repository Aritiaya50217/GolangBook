package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

const OpERROR = 5

func sendError(conn *net.UDPConn, addr *net.UDPAddr, code uint16, msg string) {
	packet := make([]byte, 4+len(msg)+1)

	binary.BigEndian.PutUint16(packet[0:2], OpERROR)
	binary.BigEndian.PutUint16(packet[2:4], code)
	copy(packet[4:], []byte(msg))
	packet[len(packet)-1] = 0

	conn.WriteToUDP(packet, addr)
	fmt.Println("ERROR sent : ", msg)

}

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":9000")
	conn, _ := net.ListenUDP("udp", addr)

	defer conn.Close()

	buf := make([]byte, 1024)

	n, clientAddr, _ := conn.ReadFromUDP(buf)
	filename := string(buf[:n])
	fmt.Println("Client requested : ", filename)

	sendError(conn, clientAddr, 1, "File not found. ")
}
