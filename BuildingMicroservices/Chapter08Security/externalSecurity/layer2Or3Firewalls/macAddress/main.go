package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error getting interfaces: %v", err)
	}

	for _, iface := range interfaces {
		// ข้าม interface ที่ไม่มี MAC
		if iface.HardwareAddr == nil {
			continue
		}

		fmt.Printf("Interface: %s\n", iface.Name)
		fmt.Printf("  MAC Address: %s\n", iface.HardwareAddr.String())
		fmt.Printf("  Flags: %s\n\n", iface.Flags.String())
	}
}
