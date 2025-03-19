package routes

import (
	"gin-gorm/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {
	route.GET("/", controller.GetUser)
	route.POST("/", controller.CreateUser)
	route.DELETE("/:id", controller.DeleteUser)
	route.PUT("/:id", controller.UpdateUser)
}
