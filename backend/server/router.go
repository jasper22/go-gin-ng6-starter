package server

import (
    "github.com/gin-gonic/gin"
    "github.com/app8izer/go-gin-ng6-starter/backend/controllers"
    "github.com/app8izer/go-gin-ng6-starter/backend/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default() // init router with default mw (e.g. logging)

    public := r.Group("api/public")
    {
        controllers.AuthController(public)
    }

    r.GET("/api/json/:userId", controllers.GetUserByIdJson) // Todo: remove

    secure := r.Group("api/secure")
    {
        secure.Use(middlewares.AuthMiddleware())
        controllers.UserController(secure)
    }

    return r
}
