package routes

import (
	controller "go-lang-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}

//Our Auth routes accept incoming routes as gin engine and then we make our signup and login routes
//Here controllers will have the business logic to handle the routes
