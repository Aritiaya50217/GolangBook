package main

import (
	"crypto/sha1"
	"fmt"
)

func CreateHash(byteStr []byte) []byte {
	hashVal := sha1.New()
	hashVal.Write(byteStr)

	bytes := hashVal.Sum(nil)
	return bytes
}

func CreateHashMultiple(byteStr1, byteStr2 []byte) []byte {
	return xor(CreateHash(byteStr1), CreateHash(byteStr2))
}

func xor(byteStr1, byteStr2 []byte) []byte {
	xorbytes := make([]byte, len(byteStr1))
	for i := 0; i < len(byteStr1); i++ {
		xorbytes[i] = byteStr1[i] ^ byteStr2[i]
	}
	return xorbytes
}

func main() {
	bytes := CreateHashMultiple([]byte("Check"), []byte("Hash"))
	fmt.Printf("%x\n", bytes)
}
