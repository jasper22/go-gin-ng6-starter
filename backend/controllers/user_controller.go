package controllers

import (
    "net/http"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/app8izer/go-gin-ng6-starter/backend/services"
    "github.com/app8izer/go-gin-ng6-starter/backend/models"
)

var userService = services.GetUserServiceInstance()

func UserController(r *gin.RouterGroup) {
    user := r.Group("user")
    {
        user.POST("/", CreateUser)
        user.GET("/:userId", GetUserById)
        user.GET("/", GetAllUsers)
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

func GetAllUsers(c *gin.Context) {
    res, err := json.Marshal(userService.GetAll())
    if err == nil {
        c.JSON(http.StatusOK, string(res))
        return
    }
    c.JSON(http.StatusBadRequest, gin.H{"error": "Error"})
}

func CreateUser(c *gin.Context) {
    user := models.User{}
    if err := c.ShouldBind(&user); err == nil {
        res, err := json.Marshal(userService.Create(&user))
        if err == nil {
            c.JSON(http.StatusOK, string(res))
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    }
}
