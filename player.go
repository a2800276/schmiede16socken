package socken

import "container/list"

type Game struct {
	Players *list.List
	Cards   *list.List
	Round   int
	View    View
}
type Player struct {
	Name   string
	Number int
	Card   *Card
	Score  int
	Game   *Game
}

var playerCounter = 0

func NewGame() *Game {
	game := new(Game)
	game.Players = list.New()
	game.Cards = list.New()
	for i := 0; i != NUM_CARDS; i++ {
		game.Cards.PushBack(GetCard(i))
	}
	return game
}

// initialize a new Player ...
func (g *Game) AddPlayer(name string) *Player {
	defer func() {
		playerCounter++
	}()

	player := &Player{
		name,
		playerCounter,
		g.GetCard(),
		0,
		g,
	}

	g.Players.PushBack(player)
	return player
}

func (g *Game) RemovePlayer(player *Player) {
	elem := findElement(player, g.Players)
	if nil == elem {
		return
	}
	g.Players.Remove(elem)
}

func (g *Game) GetCard() *Card {
	elem := g.Cards.Front()
	if elem == nil {
		return nil
	}
	card := elem.Value
	g.Cards.Remove(elem)
	return card.(*Card)
}

func (g *Game) ReturnCard(card *Card) {
	if nil != findElement(card, g.Cards) {
		return
	}
	g.Cards.PushBack(card)
}

func (p *Player) Guess(sym Symbol) {

}

func findElement(value interface{}, l *list.List) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return e
		}
	}
	return nil
}
