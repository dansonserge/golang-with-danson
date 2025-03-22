package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs the request details
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// Process request
		c.Next()

		// After request
		duration := time.Since(startTime)

		log.Printf("Danson serge Method: %s | Path: %s | Status: %d | Duration: %v\n",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)

	}
}

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v\n", err)
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()
		c.Next()
	}
}
