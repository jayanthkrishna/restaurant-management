package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/restaurant-management-golang/controllers"
)

func MenuRoutes(router *gin.Engine) {

	router.GET("/menus", controllers.GetMenus())
	router.GET("/menus/:menu_id", controllers.GetMenu())
	router.POST("/menus", controllers.CreateMenu())
	router.PATCH("/menus/:menus_id", controllers.UpdateMenu())

}
