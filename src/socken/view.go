package socken

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

var TheView = NewView()

type View interface {
	//AddPlayer(string, *websocket.Conn)
	Broadcast(msg string)
	CorrectGuess(guess Symbol, p *Player)
	NewPlayerCard(*Player)
	BoardCard()
	IncorrectGuess(player *Player)
}

type DummyView struct{}

func (d DummyView) Broadcast(_ string)                   {}
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
	view.Game.View = view
	view.PlayerSockets = make(map[string]*websocket.Conn)
	return &view
}

func (v WebsocketView) AddPlayer(name string, c *websocket.Conn) {
	player := v.Game.AddPlayer(name)

	println(player.Game)
	println(player.Game.View)
	println(&v)

	v.PlayerSockets[name] = c
	v.NewPlayerCard(player)

	data := fmt.Sprintf("\"%s\"", player.Name)
	sendSharedSocket("addPlayer", data)
	v.Broadcast(fmt.Sprintf("Welcome %s!", name))

}

func (v WebsocketView) Broadcast(msg string) {
	message := fmt.Sprintf("msg:%s", msg)
	sendSharedSocket("msg", fmt.Sprintf("\"%s\"", message))
	for _, s := range v.PlayerSockets {
		websocket.Message.Send(s, message)
	}
}

func (v WebsocketView) Guess(guess Symbol, c *websocket.Conn) {
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

func (v WebsocketView) CorrectGuess(guess Symbol, player *Player) {
	println("wtf, man?")
	data := fmt.Sprintf("[\"%s\", %d]", player.Name, guess)
	sendSharedSocket("correctGuess", data)
	v.UpdatePlayer(player)
}
func (v WebsocketView) IncorrectGuess(player *Player) {
	data := fmt.Sprintf("\"%s\"", player.Name)
	sendSharedSocket("incorrectGuess", data)
	v.UpdatePlayer(player)
}

func (v WebsocketView) UpdatePlayer(player *Player) {
	data := fmt.Sprintf("[\"%s\", %d]", player.Name, player.Score)
	sendSharedSocket("playerUpdate", data)
}
func (v WebsocketView) BoardCard() {
	data := encodeCard(v.Game.Card)
	sendSharedSocket("newCard", data)
}
func (v WebsocketView) NewPlayerCard(p *Player) {
	socket := TheView.PlayerSockets[p.Name]

	fmt.Printf("player: %s\n", encodeCard(p.Card))
	fmt.Printf("game: %s\n", encodeCard(v.Game.Card))

	websocket.Message.Send(socket, fmt.Sprintf("newCard:%s", encodeCard(p.Card)))
}

// 'data' param must be quoted if it's a string!
// messages are:
// - correctGuess [name, symbol]
// - incorrectGuess name
// - playerUpdate [name, score]
func sendSharedSocket(name string, data string) {
	message := fmt.Sprintf("{\"name\":\"%s\",\"data\":%s}", name, data)
	websocket.Message.Send(TheView.SharedSocket, message)
}

func encodeCard(card *Card) string {
	i, _ := json.Marshal(card.Symbols())
	return string(i)
}
