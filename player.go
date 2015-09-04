package main

type Player struct {
	X  int    `json:"x"`
	Y  int    `json:"y"`
	Id string `json:"id"`
}

var players map[*websocket.Conn]Player //Holds all of the players

func playerHandler(msg []byte, sender *websocket.Conn) {
	var temp Player            //Creates a new player object
	json.Unmarshal(msg, &temp) //Parses the msg into the new player object
	players[sender] = temp     //Adds/updates the player into the players map, which holds type Player
	for conn := range connections {
		if conn != sender {
			conn.WriteMessage(websocket.TextMessage, msg) //Sends the update to everybody but the player
		}
	}
	log.Println(players) //Prints out the map of all of the players
}
