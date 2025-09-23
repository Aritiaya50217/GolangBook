package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var jwtSecret []byte
var defaultPort = ":8080"

func init() {
	// หา path ของไฟล์ .env ในโฟลเดอร์เดียวกับ gateway.go
	dir, _ := os.Getwd()
	envPath := filepath.Join(dir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found at", envPath, ", using system environment")
	}

	key := os.Getenv("API_KEY")
	if key == "" {
		log.Fatal("API_KEY is not set in environment")
	}
	jwtSecret = []byte(key)

	if os.Getenv("PORT") != "" {
		defaultPort = ":" + os.Getenv("PORT")
	}
}

func authMiddleware(w http.ResponseWriter, r *http.Request) bool {
	token := r.Header.Get("Authorization")
	if token == "" || !validateToken(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}

func validateToken(tokenStr string) bool {
	tokenStr = strings.TrimSpace(tokenStr)
	if strings.HasPrefix(tokenStr, "Bearer ") {
		tokenStr = tokenStr[7:]
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		log.Println("JWT parse error:", err)
		return false
	}
	if !token.Valid {
		log.Println("JWT is invalid")
		return false
	}
	return true
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !authMiddleware(w, r) {
			return
		}
		routeHandler(w, r)
	})

	log.Println("API Gateway running on", defaultPort)
	if err := http.ListenAndServe(defaultPort, mux); err != nil {
		log.Fatal("Sever error : ", err)
	}
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case matchPath(r.URL.Path, "/users/"):
		proxyRequest(w, r, "http://localhost:9001")
	case matchPath(r.URL.Path, "/orders/"):
		proxyRequest(w, r, "http://localhost:9002")
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func matchPath(path, prefix string) bool {
	return len(path) >= len(prefix) && path[:len(prefix)] == prefix
}

func proxyRequest(w http.ResponseWriter, r *http.Request, target string) {
	targetURL, _ := url.Parse(target)
	proxyReq, _ := http.NewRequestWithContext(context.Background(), r.Method, targetURL.String()+r.URL.Path, r.Body)
	proxyReq.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		for _, val := range v {
			w.Header().Add(k, val)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
