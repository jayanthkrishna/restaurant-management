package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/restaurant-management-golang/controllers"
)
func FoodRoutes(router *gin.Engine) {	

	router.GET("/foods",controllers.GetFoods())
	router.GET("/foods/s:food_id",controllers.GetFood())
	router.POST("/foods",controllers.CreateFood())
	router.PATCH("/foods/:food_id",controllers.UpdateFood())
	

}
