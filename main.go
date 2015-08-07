package main

import (
	"flag"
	"log"
	"net/http"
	//"github.com/rs/cors"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type Thing struct {
	Typeof string `json:"type"`
}

var connections map[*websocket.Conn]bool

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Taken from gorilla's website
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	log.Println("Succesfully upgraded connection")
	connections[conn] = true

	for {
		// Blocks until a message is read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			conn.Close()
			return
		}
		var data Thing
		json.Unmarshal(msg, &data)
		if data.Typeof == "player" {
			playerHandler([]byte(data.Typeof), conn)
		}
		log.Println(data)
		log.Println(string(msg))

	}
}

func main() {
	// command line flags
	port := flag.Int("port", 5000, "port to serve on")
	flag.Parse()
	connections = make(map[*websocket.Conn]bool)
	players = make(map[*websocket.Conn]Player)

	http.HandleFunc("/ws", wsHandler)
	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
