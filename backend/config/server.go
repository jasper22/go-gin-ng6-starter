package config

import (
    "github.com/gin-gonic/gin"
)

func init(){
    cfg := GetConfig() // load global viper config
    if cfg.GetString("GIN_MODE") == "release"{
        gin.SetMode(gin.ReleaseMode)
    }
}
