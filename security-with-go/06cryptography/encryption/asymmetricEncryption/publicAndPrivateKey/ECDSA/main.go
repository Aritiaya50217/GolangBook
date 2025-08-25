package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	// สร้าง private key (ใช้ curve P256)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey
	message := []byte("Hello ECDSA")
	hash := sha256.Sum256(message)

	// sign
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("Signature: (r: %x, s: %x)\n", r, s)

	// Verify
	valid := ecdsa.Verify(publicKey, hash[:], r, s)
	fmt.Printf("Valid: %v\n", valid)
}
