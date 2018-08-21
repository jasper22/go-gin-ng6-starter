package config

import (
    "github.com/gin-gonic/gin"
)

func init() {
    cfg := GetConfig() // load global viper config
    switch cfg.GetString("GIN_MODE") {
    case "release":
        gin.SetMode(gin.ReleaseMode)
    case "test":
        gin.SetMode(gin.TestMode)
    default:
        gin.SetMode(gin.DebugMode)
    }
}
