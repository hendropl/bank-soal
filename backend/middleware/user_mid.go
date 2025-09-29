package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/helper"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}
		tokenString := tokenParts[1]

		claims, err := helper.ParseAndValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
