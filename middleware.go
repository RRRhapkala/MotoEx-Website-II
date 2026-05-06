package main

import (
	"crypto/subtle"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(adminKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if subtle.ConstantTimeCompare([]byte(adminKey), []byte(header)) != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
