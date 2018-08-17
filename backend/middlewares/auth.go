package middlewares

import (
    "github.com/gin-gonic/gin"
    "log"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("using auth mw")
        c.Next()
    }
}
