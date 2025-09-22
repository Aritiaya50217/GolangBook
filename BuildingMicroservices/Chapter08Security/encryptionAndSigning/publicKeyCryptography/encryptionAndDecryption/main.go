package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	message := []byte("Hello Public Key Cryptography!")

	// Encrypt with Public Key
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		panic(err)
	}

	// Decrypt with Private Key
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Ciphertext:", ciphertext)
	fmt.Println("Decrypted:", string(plaintext))
}
