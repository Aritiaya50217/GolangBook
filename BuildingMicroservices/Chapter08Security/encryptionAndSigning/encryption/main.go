package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// AES-GCM
func main() {
	// create key 32 bytes (AES-256)
	key := []byte("12345678901234567890123456789012")

	// ข้อความที่ต้องการเข้ารหัส
	plaintext := []byte("Hello Symmetric Encryption in Go!")

	// Encrypt
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("Ciphertext : ", encoded)

	// Decrypt
	data, _ := base64.StdEncoding.DecodeString(encoded)
	nonce, ciphertext = data[:aesGCM.NonceSize()], data[aesGCM.NonceSize():]
	plaintextOut, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted : ", string(plaintextOut))

}
