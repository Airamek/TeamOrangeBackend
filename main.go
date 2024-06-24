package main

import (
	"github.com/gin-gonic/gin"
	"main/BackendUtils/DataProvider"
)

func main() {
	var provider = DataProvider.CreateProvider("LdapUserProvider")
	provider.Init("czimer-ldap")

	// Create a new Gin router
	router := gin.Default()

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, provider.GetUsersData())
	})

	// Run the server on port 8080
	router.Run(":8080")
}
