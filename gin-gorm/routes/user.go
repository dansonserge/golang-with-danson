package routes

import (
	"gin-gorm/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {
	route.GET("/", controller.GetUsers)
	route.POST("/", controller.CreateUsers)
	route.DELETE("/:id", controller.DeleteUsers)
	route.PUT("/:id", controller.UpdateUsers)
}
