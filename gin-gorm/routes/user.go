package routes

import "github.com/gin-gonic/gin"

func UserRoute(route *gin.Engine){
 route.GET("/", func(c *gin.Context){
		c.String(200, "user route!")
	})
}