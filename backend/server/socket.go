package server

import (
    "github.com/googollee/go-socket.io"
    "log"
)

func SetupSocketServer() (socketServer *socketio.Server) {
    socketServer, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal("Failed to start socket server")
    }
    return socketServer
}
