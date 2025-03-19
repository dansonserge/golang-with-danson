package routes

import (
	"gin-gorm/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine){
	route.GET("/", controller.UserController)
}