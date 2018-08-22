package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/googollee/go-socket.io"
    "log"
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
        // emit welcome
        so.Emit("connect", gin.H{"message": "connected"})

        // join room (is only done in BE)
        so.On("join_room", func(room string) {
            so.Join(room)
            // broadcast to room with event `connected` and message `New client connected`
            socketServer.BroadcastTo(room, "connected", "New client connected")
        })

        // handler for ping messages
        so.On("ping", func(msg string) {
            log.Println("Received msg: ", msg)
        })

        // disconnect handler
        so.On("disconnection", func() {
            log.Println("Disconnected client")
        })
    })

    socketServer.On("error", func(so socketio.Socket, err error) {
        log.Printf("[ WebSocket ] Error : %v", err.Error())
    })

    socketServer.ServeHTTP(c.Writer, c.Request)
}
