package server

import (
    "go-gin-ng6-starter/backend/config"
)

func Init() {
    // setup router
    r := SetupRouter()

    // get port from env or use default 8080
    cfg := config.GetConfig()
    port := cfg.GetString("gin.port")

    // run http server
    r.Run(":" + port)

}
