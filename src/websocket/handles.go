package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func send_broadcast(message *Message) {
	// log.Println("[SENDING] Message => ", message)
	// log.Println("[SENDING] Users => ", users)

	for _, v := range users {
		if v.Topic == message.Topic || message.Topic == "all" {
			v.Conn.WriteMessage(websocket.TextMessage, []byte(message.String()))
		}
	}
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	var user UserLogged

	if err != nil {
		log.Println("[ERROR] Connection error =>  ", err.Error())
		return
	}

	iduser := r.URL.Query().Get("id")
	topic := r.URL.Query().Get("topic")

	if iduser == "" {
		conn.Close()
		return
	}

	if topic == "" {
		topic = "all"
	}

	user.Topic = topic
	user.Iduser = iduser

	user.Conn = conn
	users[iduser] = user

	log.Println("[CONNECTION] User Logged => " + iduser + " => on topic =>  " + topic)

	for {
		_, response, err := conn.ReadMessage()
		message := string(response)
		if err != nil {
			log.Println("[WARNING] => User "+iduser+" =>", err.Error())
			break
		}

		log.Println("[MESSAGE] Received message => User " + iduser)

		var target *Message
		er := json.Unmarshal([]byte(message), &target)

		if er != nil {
			log.Println("[ERROR] Read message error => User "+iduser+" =>", er.Error())
		}

		send_broadcast(target)

	}

}
