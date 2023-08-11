package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextWithTimeOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// CORS middleware function
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
