package server

import (
    "github.com/gin-gonic/gin"
    "go-gin-ng6-starter/backend/controllers"
    "go-gin-ng6-starter/backend/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default() // init router with default mw (e.g. logging)

    public := r.Group("api/public")
    {
        controllers.AuthController(public)
    }

    secure := r.Group("api/secure")
    {
        secure.Use(middlewares.AuthMiddleware())
        controllers.UserController(secure)
    }

    return r
}
