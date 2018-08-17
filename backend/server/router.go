package server

import (
    "github.com/gin-gonic/gin"
    "github.com/app8izer/go-gin-ng6-starter/backend/controllers"
    "github.com/app8izer/go-gin-ng6-starter/backend/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default() // init router with default mw (e.g. logging)

    r.GET("/api/json/:userId", controllers.GetUserByIdJson)

    // TODO: update middlewares here
    r.Use(middlewares.AuthMiddleware())

    user := r.Group("/api/user")
    {
        user.POST("/", controllers.CreateUser)
        user.GET("/:userId", controllers.GetUserById)
        user.GET("/", controllers.GetAllUsers)
    }

    return r
}
