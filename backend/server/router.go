package server

import (
    "github.com/gin-gonic/gin"
    "github.com/app8izer/go-gin-ng6-starter/backend/controllers"
    "github.com/app8izer/go-gin-ng6-starter/backend/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default() // init router with default mw (e.g. logging)

    // init authMiddleware
    authMiddleware := middlewares.GetAuthMiddleware()

    public := r.Group("api/public")
    {
        controllers.AuthController(authMiddleware, public)
    }

    secure := r.Group("api/secure")
    {

        secure.Use(authMiddleware.MiddlewareFunc())
        controllers.UserController(secure)
    }

    admin := r.Group("api/admin")
    {
        admin.Use(authMiddleware.MiddlewareFunc())
        admin.Use(authMiddleware.AdminMiddleware)
        controllers.AdminController(admin)

    }

    return r
}
