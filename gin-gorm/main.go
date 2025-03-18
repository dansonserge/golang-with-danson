package main

import (
	"gin-gorm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.UserRoute(router)

	router.Run(":8001")
}