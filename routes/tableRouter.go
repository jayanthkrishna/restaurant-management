package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/restaurant-management-golang/controllers"
)

func TableRoutes(router *gin.Engine) {

	router.GET("/tables", controllers.GetTables())
	router.GET("/tables/:table_id", controllers.GetTable())
	router.POST("/tables", controllers.CreateTable())
	router.PATCH("/tables/:table_id", controllers.UpdateTable())

}
