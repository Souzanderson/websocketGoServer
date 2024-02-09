package websocket

import "github.com/gorilla/websocket"

type UserLogged struct {
	Iduser string
	Conn   *websocket.Conn
	Topic  string
}
