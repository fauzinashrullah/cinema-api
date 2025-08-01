package middleware

import (
	"net/http"
	"strings"

	"github.com/fauzinashrullah/cinema-api/utils"
	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
            return
        }
        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := utils.ParseJWT(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            return
        }
        c.Set("userID", claims.UserID)
        c.Next()
    }
}
