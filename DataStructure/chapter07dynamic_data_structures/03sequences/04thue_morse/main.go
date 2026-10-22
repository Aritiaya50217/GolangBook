package main

import (
	"bytes"
	"fmt"
)

func ThueMorseSequence(buffer *bytes.Buffer) {
	for b, currLength, currBytes := 0, buffer.Len(), buffer.Bytes(); b < currLength; b++ {
		if currBytes[b] == '1' {
			buffer.WriteByte('0')
		} else {
			buffer.WriteByte('1')
		}
	}
}

func main() {
	buffer := bytes.Buffer{}
	// initial sequece member is "0"
	buffer.WriteByte('0')
	fmt.Println(buffer.String())

	for i := 2; i <= 7; i++ {
		ThueMorseSequence(&buffer)
		fmt.Println(buffer.String())
	}
}
