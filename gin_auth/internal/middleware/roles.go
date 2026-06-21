package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// acts as a guard for admin routes
func RequireAdmin() gin.HandlerFunc {
	return func (c *gin.Context) {
		role, ok := GetRole(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		if !strings.EqualFold(role, "admin") {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Route can only be accessed by admin",
			})
			return
		}
		c.Next()
	}
}