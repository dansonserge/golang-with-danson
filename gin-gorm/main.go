package main

import (
	"gin-gorm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.UserRoute(router)

	/* router.GET("/", func(c *gin.Context){
		c.String(200, "Hello world!")
	}) */
	router.Run(":8001")
}