package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

var StaticServer http.Handler = http.FileServer(http.Dir("./assets"))

var playerSockets map[string]*websocket.Conn = make(map[string]*websocket.Conn)
var sharedSockets map[string]*websocket.Conn = make(map[string]*websocket.Conn)

// Echo the data received on the WebSocket.
func PlayerServer(ws *websocket.Conn) {
	// add to all
	var msg string
	for {
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			fmt.Printf("here2 %q\n", err)
			// remove from all
			return
		}
		fmt.Printf("msg: %s\n", msg)
		websocket.Message.Send(ws, "Hi There!")
	}
}

func SharedServer(ws *websocket.Conn) {
	var msg string
	for {
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			// waaahhh!
			return
		}
	}
}

var SharedBoard http.Handler = websocket.Handler(SharedServer)
var PlayerBoard http.Handler = websocket.Handler(PlayerServer)
