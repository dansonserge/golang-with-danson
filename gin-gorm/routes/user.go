package routes

import (
	"gin-gorm/controller"
	"gin-gorm/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type, Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	route.Use(cors.New(corsConfig))
	route.Use(middleware.LoggerMiddleware())
	route.GET("/", controller.ReGetUser)
	route.GET("/:id", controller.ReGetUserById)
	route.POST("/", controller.CreateUser)
	route.DELETE("/:id", controller.DeleteUser)
	route.PUT("/:id", controller.UpdateUser)
}
