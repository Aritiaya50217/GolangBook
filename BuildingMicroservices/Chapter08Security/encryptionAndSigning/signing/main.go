package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	// symmetric key
	encryptionKey := []byte("12345678901234567890123456789012") // 32 bytes
	// key for HMAC (sign)
	signKey := []byte("hmac-secret-key-32bytes-long--------")

	plaintext := []byte("Hello Encryption + Signing in Go!")

	// encrypt
	block, _ := aes.NewCipher(encryptionKey)
	aesgcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, aesgcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fullCipher := append(nonce, ciphertext...)

	// sign (HMAC)
	mac := hmac.New(sha256.New, signKey)
	mac.Write(fullCipher)
	signature := mac.Sum(nil)

	fmt.Println("Ciphertext (base64) : ", base64.StdEncoding.EncodeToString(fullCipher))
	fmt.Println("Signature (base64) : ", base64.StdEncoding.EncodeToString(signature))

	// verify
	mac2 := hmac.New(sha256.New, signKey)
	mac2.Write(fullCipher)
	if !hmac.Equal(signature, mac2.Sum(nil)) {
		panic("Signature invalid")
	}

	// decrypt
	nonce, ciphertext = fullCipher[:aesgcm.NonceSize()], fullCipher[aesgcm.NonceSize():]
	plaintextOut, _ := aesgcm.Open(nil, nonce, ciphertext, nil)
	fmt.Println("Decrypted : ", string(plaintextOut))

}
