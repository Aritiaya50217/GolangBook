package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encryptAES(key, plaintext []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	// สร้าง IV (initialzation Vector)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, iv, nil
}

func decryptAES(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := ciphertext[:aes.BlockSize] // ใส่ IV ไว้ตอนต้น ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}

func main() {
	key := []byte("1234567890123456") // 16 bytes = AES-128
	plaintext := []byte("Hello AES in Go!")

	// Encrypt
	ciphertext, iv, err := encryptAES(key, plaintext)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ciphertext: %x\nIV: %x\n", ciphertext, iv)

	// Decrypt
	decrypted, err := decryptAES(key, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
