package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tmpl = `
<form method="POST" action="/submit">
  <input type="hidden" name="csrf_token" value="{{.}}">
  <input type="text" name="data">
  <button type="submit">Submit</button>
</form>
`

var csrfToken string

func generateCSRFToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	csrfToken = generateCSRFToken()
	temp, _ := template.New("form").Parse(tmpl)
	temp.Execute(w, csrfToken)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Parse error", http.StatusBadRequest)
		return
	}

	token := r.Form.Get("csrf_token")
	if token != csrfToken {
		http.Error(w, "CSRF token invalid", http.StatusForbidden)
		return
	}
	data := r.Form.Get("data")
	fmt.Fprintf(w, "Form submitted successfully: %s", data)
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
