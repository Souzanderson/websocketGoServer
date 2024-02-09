package main

import (
	wsserver "gows/src/websocket"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		wsserver.RunServer("localhost", "8083")
	} else {
		wsserver.RunServer("localhost", os.Args[1])
	}
}
