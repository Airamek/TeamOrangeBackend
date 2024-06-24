package main

import (
	"github.com/gin-gonic/gin"
	"main/app/czimer-appointments"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	czimer_appointments.Routes(router)

	// Run the server on port 8080
	router.Run(":8080")
}
