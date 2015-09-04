func wsHandler(w http.ResponseWriter, r *http.Request) {
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
		} else if data.Typeof == "spell" {
			spellHandler([]byte(data.Typeof), conn)
		}
		log.Println(data)
		log.Println(string(msg))
	}
}