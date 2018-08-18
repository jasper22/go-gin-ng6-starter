package controllers

import (
    "net/http"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "go-gin-ng6-starter/backend/services"
    "go-gin-ng6-starter/backend/models"
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
    err := c.ShouldBind(&user)
    if err == nil {
        res, err := json.Marshal(userService.Create(&user))
        c.JSON(http.StatusOK, string(res))
        if err == nil {
            return
        }
    }
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
