package socken

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type View interface {
	Flash(msg string)
}

type DummyView struct{}

func (d DummyView) Flash(_ string) {}

type WebsocketView struct {
	Game          *Game
	SharedSocket  *websocket.Conn
	PlayerSockets map[string]*websocket.Conn
}

func NewView() *WebsocketView {
	view := WebsocketView{}
	view.Game = NewGame()
	view.PlayerSockets = make(map[string]*websocket.Conn)
	return &view
}

func (v *WebsocketView) AddPlayer(name string, c *websocket.Conn) {
	v.Game.AddPlayer(name)
	v.PlayerSockets[name] = c
	v.Broadcast(fmt.Sprintf("Welcome: %s!", name))
}

func (v *WebsocketView) Broadcast(msg string) {
	message := fmt.Sprintf("msg:%s", msg)
	websocket.Message.Send(v.SharedSocket, message)
	for _, s := range v.PlayerSockets {
		websocket.Message.Send(s, message)
	}
}

func (v *WebsocketView) Guess(guess Symbol, c *websocket.Conn) {
	// find name
	var playerName = ""
	for name, s := range v.PlayerSockets {
		if c == s {
			playerName = name
			break
		}
	}

	player := v.Game.GetPlayerByName(playerName)
	player.Guess(Symbol(guess))
}
