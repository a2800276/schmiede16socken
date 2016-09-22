package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"strings"
)

import (
	"socken"
)

var StaticServer http.Handler = http.FileServer(http.Dir("./assets"))

var WSView = socken.NewView()

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

		if strings.Index(msg, "Hello") == 0 {
			// Hello:<name>
			name := strings.Split(msg, ":")[1]
			WSView.AddPlayer(name, ws)
		} else {
			println("here, man")
			// guess -- parseInt
			var i int64 = -1
			i, _ = strconv.ParseInt(msg, 10, 32)
			WSView.Guess(socken.Symbol(i), ws)
		}

	}
}

func SharedServer(ws *websocket.Conn) {
	var msg string
	for {
		WSView.SharedSocket = ws
		WSView.BoardCard()
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			// waaahhh!
			return
		}
	}
}

var SharedBoard http.Handler = websocket.Handler(SharedServer)
var PlayerBoard http.Handler = websocket.Handler(PlayerServer)
