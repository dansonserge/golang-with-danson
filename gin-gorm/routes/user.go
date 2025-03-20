package routes

import (
	"gin-gorm/controller"
	"gin-gorm/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {
	route.Use(middleware.LoggerMiddleware())
	route.GET("/", controller.ReGetUser)
	route.GET("/:id", controller.ReGetUserById)
	route.POST("/", controller.CreateUser)
	route.DELETE("/:id", controller.DeleteUser)
	route.PUT("/:id", controller.UpdateUser)
}
