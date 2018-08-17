package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/appleboy/gin-jwt"
)

func AuthController(authMiddleware *jwt.GinJWTMiddleware, r *gin.RouterGroup) {
    auth := r.Group("auth")
    {
        // LoginHandler implemented in auth.go -> init() -> Author
        auth.POST("/login", authMiddleware.LoginHandler)
    }
}

/*func LoginUser(c *gin.Context) {
    log.Println("Logging in")
    // handle login
}
*/
