package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/restaurant-management-golang/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers())
	router.GET("/users/:id", controllers.GetUser())
	router.POST("/users/signup", controllers.SignUp())
	router.POST("/users/login", controllers.Login())

}
