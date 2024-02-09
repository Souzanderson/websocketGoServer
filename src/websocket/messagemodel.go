package websocket

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Topic   string      `json:"topic"`
	Payload interface{} `json:"payload"`
}

func (u Message) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf(string(b))
}
