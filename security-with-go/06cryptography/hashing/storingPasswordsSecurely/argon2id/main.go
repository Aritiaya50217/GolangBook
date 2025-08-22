package main

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func hashPassword(password string) (string, []byte) {
	salt := make([]byte, 16)
	_, _ = rand.Read(salt)

	hash := argon2.IDKey([]byte(password), salt, 3, 64*1024, 1, 32)
	encoded := base64.RawStdEncoding.EncodeToString(hash)
	return encoded, salt
}

func verifyPassword(password, encoded string, salt []byte) bool {
	hash := argon2.IDKey([]byte(password), salt, 3, 64*1024, 1, 32)
	return subtle.ConstantTimeCompare(hash, []byte(encoded)) == 1
}

func main() {
	pass := "MyS3cret!"
	hash, salt := hashPassword(pass)
	fmt.Println("Hash : ", hash)

	ok := verifyPassword(pass, hash, salt)
	fmt.Println("Password valid : ", ok)
}
