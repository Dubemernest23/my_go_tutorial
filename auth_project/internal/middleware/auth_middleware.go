package middleware

import (
	"auth_project/internal/auth"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// store auth data in context for handlers to use
const (
	ContextUserIDKey   = "auth.userID"
	ContextUserRoleKey = "auth.role"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for auth middleware
		// check if header is present
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":   false,
				"error":     "Authorization header missing",
				"timestamp": time.Now().UTC(),
			})
			return
		}
		// split bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":   false,
				"error":     "Invalid Authorization header format",
				"timestamp": time.Now().UTC(),
			})
			return
		}

		scheme := strings.TrimSpace(parts[0])
		token := strings.TrimSpace(parts[1])
		// check if scheme is Bearer
		if !strings.EqualFold(scheme, "bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":   false,
				"error":     "Authorization scheme must be Bearer",
				"timestamp": time.Now().UTC(),
			})
			return
		}

		// check for missing token
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":   false,
				"error":     "Token missing",
				"timestamp": time.Now().UTC(),
			})
			return
		}

		claims, err := auth.ParseToken(jwtSecret, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":   false,
				"error":     "Invalid token",
				"timestamp": time.Now().UTC(),
			})
			return
		}

		// store auth data in context for handlers to use
		c.Set(ContextUserIDKey, claims.Subject)
		c.Set(ContextUserRoleKey, claims.Role)

		// call next handler
		c.Next()
	}

}

func GetUserID(c *gin.Context) (string, bool) {
	res, exists := c.Get(ContextUserIDKey)
	if !exists {
		return "", false
	}
	return res.(string), true
}

func GetUserRole(c *gin.Context) (string, bool) {
	res, exists := c.Get(ContextUserRoleKey)
	if !exists {
		return "", false
	}
	return res.(string), true
}
