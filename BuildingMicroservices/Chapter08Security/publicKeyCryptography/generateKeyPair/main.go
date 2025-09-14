package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	// generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	// Export provate key as PEM
	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privPem := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: privBytes,
	})

	// Export public key as PEM
	pubBytes := x509.MarshalPKCS1PublicKey(publicKey)
	pubPem := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY", Bytes: pubBytes,
	})
	fmt.Println("Private Key:\n", string(privPem))
	fmt.Println("Public Key:\n", string(pubPem))
}
