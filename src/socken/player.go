package socken

import "container/list"

// Game can:
// - be created (NewGame)
// - manage cards (GetCard, ReturnCard)
// = manage players (AddPlayer, RemovePlayer)
type Game struct {
	Players *list.List
	Cards   *list.List
	Card    *Card
	Round   int
	View    View
}

// Player can:
// - be created (Game.AddPlayer)
// - die (Game.RemovePlayer) -- tbd periodic task to remove inactive players
// - guess (Player.Guess)
type Player struct {
	Name   string
	Number int
	Card   *Card
	Score  int
	Game   *Game
}

func NewGame() *Game {
	game := new(Game)
	game.Players = list.New()
	game.Cards = list.New()
	for i := 0; i != NUM_CARDS; i++ {
		game.Cards.PushBack(GetCard(i))
	}
	game.Card = game.GetCard()
	game.View = DummyView{}
	return game
}

// assign uuid to player to identify them on the wire.
var playerCounter = 0

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
	g.ReturnCard(player.Card)
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

func (g *Game) GetPlayerByName(name string) *Player {
	for e := g.Players.Front(); e != nil; e = e.Next() {
		player := e.Value.(*Player)
		if player.Name == name {
			return player
		}
	}
	return nil
}
func (g *Game) GetPlayerById(id int) *Player {
	for e := g.Players.Front(); e != nil; e = e.Next() {
		player := e.Value.(*Player)
		if player.Number == id {
			return player
		}
	}
	return nil

}

func (p *Player) Guess(guess Symbol) bool {
	match := p.Card.Match(p.Game.Card)
	if match == guess {
		p.Score += 1
		// view: flash score
		p.Game.View.CorrectGuess(guess, p)

		p.Game.ReturnCard(p.Card)
		p.Card = p.Game.Card
		p.Game.Card = p.Game.GetCard()
		// view: new card
		// playerView: new card

		p.Game.View.NewPlayerCard(p)
		p.Game.View.BoardCard()

		return true
	} else {
		p.Score -= 1
		p.Game.View.IncorrectGuess(p)

		// playerView : moep.
		return false
	}
}

func findElement(value interface{}, l *list.List) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return e
		}
	}
	return nil
}
