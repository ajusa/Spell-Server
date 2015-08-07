package main

type Player struct {
	X  int    `json:"x"`
	Y  int    `json:"y"`
	Id string `json:"id"`
}

var players map[*websocket.Conn]Player

func playerHandler(msg []byte, sender *websocket.Conn) {
	var temp Player
	json.Unmarshal(msg, &temp)
	players[sender] = temp
	for conn := range connections {
		if conn != sender {
			conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
	log.Println(players)
}
