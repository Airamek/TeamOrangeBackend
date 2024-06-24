package auth

import (
	"github.com/gin-gonic/gin"
	"main/BackendUtils/DataProvider"
	"main/BackendUtils/users"
	"net/http"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(provider DataProvider.UserDataProvider, c *gin.Context) *users.Session {

	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		session := CheckAuth(c)
		if session != nil {
			c.JSON(200, gin.H{"info": "Already logged in!"})
			return session
		}
	}

	credentials := new(LoginData)
	err := c.ShouldBindJSON(credentials)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	user := provider.AuthUser(credentials.Username, credentials.Password)
	println(credentials.Username)
	println(credentials.Password)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return nil
	}
	return users.AddSession(&user)
}
