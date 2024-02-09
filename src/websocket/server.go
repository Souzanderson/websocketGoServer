package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var users = make(map[string]UserLogged, 0)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func RunServer(endpoint string, port string) {
	log.Println("[INFO] 🚀🚀🚀 Listening server in  =>", endpoint+":"+port)
	http.HandleFunc("/", WsHandler)
	http.ListenAndServe(endpoint+":"+port, nil)
}
