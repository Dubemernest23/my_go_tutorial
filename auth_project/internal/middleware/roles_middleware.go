package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices" // or "slices" if using Go 1.21+
)

func forbiddenResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"success":   false,
		"error":     message,
		"timestamp": time.Now().UTC(),
	})
}

func RolesMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get(ContextUserRoleKey)
		if !exists {
			forbiddenResponse(c, "User role not found in context")
			return
		}

		role, ok := roleValue.(string)
		if !ok {
			forbiddenResponse(c, "Invalid user role type in context")
			return
		}

		if !slices.Contains(allowedRoles, role) {
			forbiddenResponse(c, "Access denied: insufficient permissions")
			return
		}

		c.Next()
	}
}

// // forbiddenResponse is a helper to avoid repeating the error response shape.
// func forbiddenResponse(c *gin.Context, message string) {
// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 		"success":   false,
// 		"error":     message,
// 		"timestamp": time.Now().UTC(),
// 	})
// }

// func RolesMiddleware(allowedRoles ...string) gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		roleValue, exists := c.Get(ContextUserRoleKey)
// 		// roleValue, exists := GetUserRole(c)
// 		if !exists {
// 			forbiddenResponse(c, "User role not found in context")
// 			return
// 		}
// 		// if !strings.EqualFold(roleValue, "admin"){
// 		// 	forbiddenResponse(c, "Access denied: admin role required")
// 		// 	return
// 		// }

// 		role, ok := roleValue.(string)
// 		if !ok {
// 			forbiddenResponse(c, "Invalid user role type in context")
// 			return
// 		}

// 		for _, allowed := range allowedRoles {
// 			if role == allowed {
// 				c.Next()
// 				return
// 			}
// 		}

// 		forbiddenResponse(c, "Access denied: insufficient permissions")
// 	}
// }
