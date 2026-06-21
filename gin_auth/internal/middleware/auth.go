package middleware

import (
	"gin_auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserId = "auth.userId"
	ctxRole = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func (c *gin.Context){
		// get auth header -> eg Authorization: Bearer(Token)
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization token",
			})
			return
		}

		seperated_string := strings.SplitN(authHeader, " ", 2) //split into 2
		if len(seperated_string) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization Format",
			})
		}

		auth_scheme := strings.TrimSpace(seperated_string[0]) //check if first part of schema is Bearer
		if !strings.EqualFold(auth_scheme, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Auth scheme must start with Bearer",
			})
			return
		}

		tokenString := strings.TrimSpace(seperated_string[1]) //extract the second part, which is the token

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing token",
			})
			return
		}

		claims, err := auth.ParseAndValidateToken(jwtSecret, tokenString) //parse and validate token
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		// setting user information in request context
		c.Set(ctxUserId, claims.Subject)
		c.Set(ctxRole, claims.Role)
		c.Next()

	}
}

func GetUserId(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxUserId)
	if !ok {
		return "", false
	}

	userID, ok := res.(string)
	return userID, ok
}

func GetRole(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxRole)
	if !ok {
		return "", false
	}

	role, ok := res.(string)
	return role, ok
}