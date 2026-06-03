package main

import "time"

var cache [][]byte

func main() {
	for {
		data := make([]byte, 1024*1024)
		cache = append(cache, data)

		time.Sleep(5 * time.Second)
		break
	}
}
