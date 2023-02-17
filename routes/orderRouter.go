package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/restaurant-management-golang/controllers"
)

func OrderRoutes(router *gin.Engine) {

	router.GET("/orders", controllers.GetOrders())
	router.GET("/orders/:order_id", controllers.GetOrder())
	router.POST("/orders", controllers.CreateOrder())
	router.PATCH("/orders/:order_id", controllers.UpdateOrder())

}
