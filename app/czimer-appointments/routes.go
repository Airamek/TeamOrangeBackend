package czimer_appointments

import (
	"github.com/gin-gonic/gin"
	"main/BackendUtils/DataProvider"
	"main/BackendUtils/auth"
)

func Routes(router *gin.Engine) {
	var provider = DataProvider.CreateProviderUser("LdapUserProvider")
	provider.Init("czimer-ldap")

	router.GET("/api/v1/users", func(c *gin.Context) {
		//session := auth.CheckAuth(c)
		//if session == nil {
		//	return
		//}
		c.JSON(200, provider.GetUsersData())
	})

	router.POST("/api/v1/login", func(c *gin.Context) {

		session := auth.Login(provider, c)
		if session == nil {
			return
		}
		c.JSON(200, gin.H{"Token": session.Token.Token})
	})
}
