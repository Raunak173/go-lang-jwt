package main

import (
	"os"

	"github.com/Raunak173/go-lang-jwt/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// First we will set the Port where to run our server

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	//Next we will create our Router

	router := gin.New()
	router.Use(gin.Logger()) //This logs api in the console

	//Then we will just use our routes

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	//Example route, where 200 shows API status code and H sets the message in the Header

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
