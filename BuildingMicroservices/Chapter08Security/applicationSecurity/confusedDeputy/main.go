package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	CanRead  bool   `json:"can_read"`
	jwt.StandardClaims
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid || !claims.CanRead {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	file := r.URL.Query().Get("file")
	if !strings.HasPrefix(file, "/safe-data/") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	data, err := os.ReadFile(file)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	w.Write(data)
}

func main() {
	http.HandleFunc("/download", downloadHandler)
	http.ListenAndServe(":8080", nil)
}
