package server

import (
    "github.com/app8izer/go-gin-ng6-starter/backend/config"
    "github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
    r = SetupRouter()
}

func StartServer() {
    // get config
    cfg := config.GetConfig()
    port := cfg.GetString("PORT")

    r.Run(":" + port)
}
