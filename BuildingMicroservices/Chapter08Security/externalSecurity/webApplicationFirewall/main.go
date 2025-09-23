package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// ตัวอย่าง regex pattern สำหรับตรวจสอบ SQL Injection / XSS
var (
	sqlInjectionPattern = regexp.MustCompile(`(?i)(union\s+select|drop\s+table|--|;|')`)
	xssPattern          = regexp.MustCompile(`(?i)(<script>|javascript:)`)
)

func wafMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ตรวจสอบ query string , body , header
		if sqlInjectionPattern.MatchString(r.URL.RawQuery) || xssPattern.MatchString(r.URL.RawQuery) {
			http.Error(w, "Blocked by WAF: malicious query", http.StatusForbidden)
			log.Printf("Blocked request: SQLi/XSS in query from %s", r.RemoteAddr)
			return
		}

		// ตรวจสอบ User-Agent (block bot)
		if strings.Contains(strings.ToLower(r.UserAgent()), "sqlmap") {
			http.Error(w, "Blocked by WAF : automated scanner", http.StatusForbidden)
			log.Printf("Blocked bot: %s", r.UserAgent())
			return
		}

		// ถ้าไม่มี pattern ที่ผิดปกติจะผ่าน
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello , this request password WAF check")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	log.Println("Starting WAF-protected server on : 8080")
	if err := http.ListenAndServe(":8080", wafMiddleware(mux)); err != nil {
		log.Fatal(err)
	}

}
