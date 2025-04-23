package routes

import (
	"user_manager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateUserController)
	router.GET("/users", controllers.GetUsersController)
	router.GET("/users/:id", controllers.GetUserByIDController)
}
