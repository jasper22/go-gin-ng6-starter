package server

import (
    "github.com/app8izer/go-gin-ng6-starter/backend/config"
    "github.com/gin-gonic/gin"
    "github.com/googollee/go-socket.io"
)

var Router *gin.Engine
var SocketServer *socketio.Server

func init() {
    SocketServer = SetupSocketServer()
    Router = SetupRouter()
}

func StartServer() {
    // get config
    cfg := config.GetConfig()
    port := cfg.GetString("PORT")

    Router.Run(":" + port)
}
