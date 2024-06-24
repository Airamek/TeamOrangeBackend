package auth

import (
	"github.com/gin-gonic/gin"
	"main/BackendUtils/users"
	"net/http"
	"strings"
)

func CheckAuth(c *gin.Context) *users.Session {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil
	}

	tokenString := authToken[1]

	session := users.CheckSession(tokenString)
	if session == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil
	}

	return session
}
