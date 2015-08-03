package main
import (
	"flag"
    "log"
    "net/http"
    //"github.com/rs/cors"
    "github.com/gorilla/websocket"
    "fmt"
    "encoding/json"
)

type Player struct {
    X   int      `json:"x"`
    Y 	int 	 `json:"y"`
    Id	string	 `json:"id"`	
}
type Arrow struct {
    X   int      `json:"x"`
    Y 	int 	 `json:"y"`
    Damage	int	 `json:"damage"` 	
}
type Response struct {
    typeof	string	 `json:"type"`	
    //data 	string   `json:"data"`		
}
var connections map[*websocket.Conn]bool
var players map[*websocket.Conn]Player

func sendAll(msg []byte, sender *websocket.Conn) {
	var temp Player
	json.Unmarshal(msg, &temp)
	players[sender] = temp
	for conn := range connections {
		if conn != sender{
		conn.WriteMessage(websocket.TextMessage, msg); 
	}
	}
		log.Println(players)
}

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
		var data Response
		json.Unmarshal(msg, &data)
		if data.typeof == "player"{
			//sendAll([]byte(data.data), conn)
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

	// handle all requests by serving a file of the same name
	http.HandleFunc("/ws", wsHandler)

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
