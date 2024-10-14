package routes

import (
	controller "github.com/Raunak173/go-lang-jwt/controllers"
	"github.com/Raunak173/go-lang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	//First protect the routes using middleware
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
