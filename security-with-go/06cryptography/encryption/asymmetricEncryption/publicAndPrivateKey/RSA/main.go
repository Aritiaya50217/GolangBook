package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// สร้าง key RSA (2048 bit)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	message := []byte("Hello RSA!")

	// Encrypt ด้วย public key
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Decrypt ด้วย private key
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext)

}
