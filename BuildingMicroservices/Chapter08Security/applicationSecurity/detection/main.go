package main

import (
	"encoding/base64"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	limiterStore = make(map[string]*rate.Limiter)
	mu           sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := limiterStore[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 5) // 1 req/sec , burst 5
		limiterStore[ip] = limiter
	}
	return limiter
}

func generateCSRToken() string {
	token := make([]byte, 32)
	rand.Read(token)
	return base64.URLEncoding.EncodeToString(token)
}

func DetectionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		// rate limiting (DDoS detection)
		if !getLimiter(ip).Allow() {
			log.Printf("[ALERT] High traffic from %s", ip)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many Request",
			})
			return
		}

		// Suspicious Input Detection
		for key, vals := range c.Request.URL.Query() {
			for _, v := range vals {
				if detectSQLi(v) {
					log.Printf("[ALERT] Possible SQL Injection from %s: param=%s value=%s", ip, key, v)
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Suspicious input"})
					return
				}
			}
		}

		// CSRF Token check (สำหรับ POST/PUT/DELETE)
		if c.Request.Method != http.MethodGet {
			token := c.GetHeader("X-CSRF-Token")
			cookie, err := c.Cookie("csrf_token")
			if err != nil || token != cookie {
				log.Printf("[ALERT] Invalid CSRF token from %s", ip)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error": "Invalid CSRF token",
				})
				return
			}
		}

		// Panic Recovery
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[ALERT] Panic detected: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
			}
		}()
		c.Next()
	}
}

func detectSQLi(input string) bool {
	suspicious := []string{"--", ";--", "/*", "*/", "@@", "xp_", " or ", " and "}
	lower := strings.ToLower(input)
	for _, s := range suspicious {
		if strings.Contains(lower, s) {
			return true
		}
	}
	return false
}
func main() {
	r := gin.Default()
	r.Use(DetectionMiddleware())

	// router
	r.GET("/token", func(c *gin.Context) {
		token := generateCSRToken()
		c.SetCookie("csrf_token", token, 3600, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"csrf_token": token})
	})

	r.POST("/secure-action", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Action executed safely!"})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Server started on : 8080")
	srv.ListenAndServe()
}
