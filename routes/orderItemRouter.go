package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/restaurant-management-golang/controllers"
)

func OrderItemRoutes(router *gin.Engine) {

	router.GET("/orderItems", controllers.GetOrderItems())
	router.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	router.GET("orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	router.POST("/orderItems", controllers.CreateOrderItem())
	router.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())

}
