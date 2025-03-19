package main

import (
	"gin-gorm/config"
	"gin-gorm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8001")
}
