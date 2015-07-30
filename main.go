package main

import (
    "log"
    "net/http"
    "github.com/rs/cors"
    "github.com/googollee/go-socket.io"
)
func main() {
    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        so.Join("game")
        so.Emit("id",'so.Id())
        log.Println("on connection")
        so.On("player", func(msg string) {
            log.Println(msg)
            server.BroadcastTo("game","player", msg)
        })
        so.On("death", func(msg string) {
            server.BroadcastTo("game","death", msg)
        })
        so.On("spell", func(msg string) {
            server.BroadcastTo("game","spell", msg)
        })
        so.On("disconnection", func() {
            log.Println("on disconnect")
            server.BroadcastTo("game","death", so.Id())
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })
    c := cors.New(cors.Options{AllowedOrigins: []string{"*"},AllowCredentials: true,})
    handler := c.Handler(server)
    http.Handle("/socket.io/", handler)
    log.Println("Serving at localhost:5000...")
    log.Fatal(http.ListenAndServe(":5000", nil))
}
