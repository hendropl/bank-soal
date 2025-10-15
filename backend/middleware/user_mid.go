package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/helper"
	"latih.in-be/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		var tokenString string
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
				tokenString = parts[1]
			}
		}

		if tokenString != "" {
			claims, err := helper.ParseAndValidateToken(tokenString)
			if err == nil {
				c.Set("user", claims)
				c.Next()
				return
			}
		}

		refreshToken, err := c.Cookie("refresh_token")
		if err != nil || refreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateRefreshToken(refreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
			c.Abort()
			return
		}

		user := &model.User{Id: userId}
		fmt.Println("DEBUG user id:", user.Id)
		newAccessToken, err := helper.GenerateAccessToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate new access token"})
			c.Abort()
			return
		}

		helper.Write(c, newAccessToken)

		claims, _ := helper.ParseAndValidateToken(newAccessToken)
		c.Set("user", claims)

		c.Next()
	}
}
