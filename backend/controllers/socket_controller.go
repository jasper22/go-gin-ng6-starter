package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/googollee/go-socket.io"
    "fmt"
)

var socketServer *socketio.Server

func SocketController(r *gin.RouterGroup, s *socketio.Server) {
    socketServer = s
    sock := r.Group("socket-io")
    {
        sock.POST("/", socketHandler)
        sock.GET("/", socketHandler)
        sock.Handle("WS", "/", socketHandler)
        sock.Handle("WSS", "/", socketHandler)
    }
}

func socketHandler(c *gin.Context) {
    socketServer.On("connection", func(so socketio.Socket) {
        fmt.Println("on connection")

        so.Emit("news", gin.H{"message": "news"})

        so.Join("chat")

        so.On("chat message", func(msg string) {
            fmt.Println("emit:", so.Emit("chat message", msg))
            so.BroadcastTo("chat", "chat message", msg)
        })
        so.On("disconnection", func() {
            fmt.Println("on disconnect")
        })
    })

    socketServer.On("error", func(so socketio.Socket, err error) {
        fmt.Printf("[ WebSocket ] Error : %v", err.Error())
    })

    socketServer.ServeHTTP(c.Writer, c.Request)
}
