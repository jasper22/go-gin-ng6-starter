package controllers

import (
    "github.com/gin-gonic/gin"
    "log"
)

func AuthController(r *gin.RouterGroup) {
    auth := r.Group("auth")
    {
        auth.POST("/login", LoginUser)
    }
}

func LoginUser(c *gin.Context) {
    log.Println("Logging in")
    // handle login
}
