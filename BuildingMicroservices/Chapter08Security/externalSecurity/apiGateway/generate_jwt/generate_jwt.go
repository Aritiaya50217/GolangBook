package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func main() {
	// หา path ของไฟล์ .env ในโฟลเดอร์เดียวกับ gateway.go
	dir, _ := os.Getwd()
	envPath := filepath.Join(dir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found at", envPath, ", using system environment")
	}

	key := os.Getenv("API_KEY")
	secret := []byte(key) // ต้องตรงกับ API_KEY
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "123456",
		"exp":  time.Now().Add(time.Hour).Unix(),
		"role": "admin",
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatal("Error signing token:", err)
	}
	fmt.Println("JWT:", tokenString)
}
