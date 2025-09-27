package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(con *gin.Context) {
		token := con.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			con.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			con.Abort()
			return
		}
		con.Next()
	}
}
