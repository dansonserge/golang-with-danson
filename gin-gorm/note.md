While Gin's minimalist approach is efficient for smaller applications, you can adopt a more Spring-like structure by introducing service and repository layers. Here's how you can refactor your current Gin app to make it more modular:

controllers
'''
package controller

import (
"gin-gorm/service"
"github.com/gin-gonic/gin"
"net/http"
)

func CreateUser(c \*gin.Context) {
var user service.UserRequest
if err := c.ShouldBindJSON(&user); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
return
}

    createdUser, err := service.CreateUser(user)
    if err != nil {
    	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
    	return
    }
    c.JSON(http.StatusCreated, createdUser)

}
'''
