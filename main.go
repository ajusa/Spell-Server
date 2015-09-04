package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Thing struct {
	Typeof string `json:"type"`
}

var connections map[*websocket.Conn]bool

func main() {

	port := flag.Int("port", 5000, "port to serve on")
	flag.Parse()
	connections = make(map[*websocket.Conn]bool)
	players = make(map[*websocket.Conn]Player)
	http.HandleFunc("/ws", wsHandler)
	log.Printf("Running on port %d\n", *port)
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
