package main

import (
	"fmt"
	"net"
)

func main() {
	names, err := net.LookupAddr("8.8.8.8")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(names)
}
