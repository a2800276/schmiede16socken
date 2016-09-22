package socken

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

type View interface {
	Flash(msg string)
	CorrectGuess(guess Symbol, p *Player)
	NewPlayerCard(*Player)
	BoardCard()
	IncorrectGuess(player *Player)
}

type DummyView struct{}

func (d DummyView) Flash(_ string)                       {}
func (d DummyView) CorrectGuess(guess Symbol, p *Player) {}
func (d DummyView) NewPlayerCard(p *Player)              {}
func (d DummyView) BoardCard()                           {}
func (d DummyView) IncorrectGuess(player *Player)        {}

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
	player := v.Game.AddPlayer(name)
	v.NewPlayerCard(player)

	v.PlayerSockets[name] = c

	v.Broadcast(fmt.Sprintf("Welcome: %s!", name))

}

func (v *WebsocketView) Broadcast(msg string) {
	message := fmt.Sprintf("msg:%s", msg)
	v.sendSharedSocket("msg", fmt.Sprintf("\"%s\"", message))
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

func (v *WebsocketView) CorrectGuess(guess Symbol, player *Player) {
	data := fmt.Sprintf("[\"%s\", %d]", player.Name, guess)
	v.sendSharedSocket("correctGuess", data)
	v.UpdatePlayer(player)
}
func (v *WebsocketView) IncorrectGuess(player *Player) {
	data := fmt.Sprintf("\"%s\"", player.Name)
	v.sendSharedSocket("incorrectGuess", data)
	v.UpdatePlayer(player)
}

func (v *WebsocketView) UpdatePlayer(player *Player) {
	data := fmt.Sprintf("[\"%s\", %d]", player.Name, player.Score)
	v.sendSharedSocket("playerUpdate", data)
}
func (v *WebsocketView) BoardCard() {
	data := encodeCard(v.Game.Card)
	v.sendSharedSocket("newCard", data)
}
func (v *WebsocketView) NewPlayerCard(p *Player) {
	socket := v.PlayerSockets[p.Name]
	websocket.Message.Send(socket, fmt.Sprintf("newCard:%s", encodeCard(p.Card)))
}

// 'data' param must be quoted if it's a string!
// messages are:
// - correctGuess [name, symbol]
// - incorrectGuess name
// - playerUpdate [name, score]
func (v *WebsocketView) sendSharedSocket(name string, data string) {
	message := fmt.Sprintf("{\"name\":\"%s\",\"data\":%s}", name, data)
	websocket.Message.Send(v.SharedSocket, message)
}

func encodeCard(card *Card) string {
	i, _ := json.Marshal(card.Symbols())
	return string(i)
}
