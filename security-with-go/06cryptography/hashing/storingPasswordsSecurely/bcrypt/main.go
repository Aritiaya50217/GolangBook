package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "MyS3cret!"

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hash : ", string(hashed))

	// check password
	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		log.Fatal("Password incorrect.")
	}
	fmt.Println("Password match.")
}
