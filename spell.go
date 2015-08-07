package main

type Spell struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Damage int `json:"damage"`
}

func spellHandler(msg []byte, sender *websocket.Conn) {

}
