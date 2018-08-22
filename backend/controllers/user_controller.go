package controllers

import (
    "net/http"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/app8izer/go-gin-ng6-starter/backend/services"
)

var userService = services.GetUserServiceInstance()

func UserController(r *gin.RouterGroup) {
    user := r.Group("user")
    {
        user.GET("/:userId", GetUserById)
    }
}

func GetUserById(c *gin.Context) {
    userId := c.Param("userId")
    if userId != "" {
        res, err := json.Marshal(userService.GetById(userId))
        if err == nil {
            c.JSON(http.StatusOK, string(res))
            return
        }
    }
    c.JSON(http.StatusBadRequest, gin.H{"error": "Error"})
}

