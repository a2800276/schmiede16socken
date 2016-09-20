package socken

import "testing"

func TestMatch(t *testing.T) {
	for i := MIN_CARD; i <= MAX_CARD; i++ {
		//_ := GetCard(i)
		thisCard := GetCard(i)

		for j := MIN_CARD; j <= MAX_CARD; j++ {
			otherCard := GetCard(j)
			if i == j {
				if (otherCard.Match(thisCard) != SAME_CARD) || (thisCard.Match(otherCard) != SAME_CARD) {
					t.Error("comparing card to itself ...")
				}
			} else {
				match := otherCard.Match(thisCard)
				match2 := thisCard.Match(otherCard)
				if match != match2 {
					t.Error("comparing cards yields different results ...")
				}

				if !contains(otherCard.Symbols(), match) ||
					!contains(otherCard.Symbols(), match2) ||
					!contains(thisCard.Symbols(), match) ||
					!contains(thisCard.Symbols(), match2) {
					t.Error("match not contained in Symbols, %d, %d", i, j)
				}
			}
		}

	}
}

func TestNewGame(t *testing.T) {
	game := NewGame()
	if game.Cards.Len() != NUM_CARDS-1 {
		t.Error("incorrect number of cards in new game")
	}
	if game.Players.Len() != 0 {
		t.Error("non zero number of players in new game")
	}

	if game.Card.Match(game.Card) != SAME_CARD {
		t.Error("no card selected!?")
	}
}

func TestAddPlayer(t *testing.T) {
	game := NewGame()
	game.AddPlayer("Spongebob")
	game.AddPlayer("Patrick")
	if game.Cards.Len() != NUM_CARDS-3 {
		t.Error("incorrect number of cards in game w/ new player")
	}

	if game.Players.Len() != 2 {
		t.Error("incorrect number of players in game w/ new player")
	}
}

func TestRemovePlayer(t *testing.T) {
	game := NewGame()
	player1 := game.AddPlayer("Spongebob")
	player2 := game.AddPlayer("Patrick")

	game.RemovePlayer(player1)
	game.RemovePlayer(player2)

	if game.Cards.Len() != NUM_CARDS-1 {
		t.Error("incorrect number of cards in game w/ removed player")
	}

	if game.Players.Len() != 0 {
		t.Error("incorrect number of players in game w/ removed player")
	}
}

func TestScoreIncorrect(t *testing.T) {
	game := NewGame()
	player1 := game.AddPlayer("Spongebob")

	wrong := player1.Card.Match(game.Card) + 1
	for i := 0; i != 1000; i++ {
		player1.Guess(wrong)
	}
	if player1.Score != -1000 {
		t.Error("incorrect negative score")
	}
}

func TestScoreCorrect(t *testing.T) {
	game := NewGame()
	player1 := game.AddPlayer("Spongebob")
	player2 := game.AddPlayer("Patrick")

	for i := 0; i != 1000; i++ {
		tmpCard := game.Card

		var player *Player
		if i%2 == 0 {
			player = player1
		} else {
			player = player2
		}
		correct := player.Card.Match(game.Card)

		player.Guess(correct)
		if tmpCard != player.Card {
			println(tmpCard.String())
			t.Error("Card not won by player")
		}
	}
	if player1.Score != 500 {
		t.Error("incorrect score")
	}
}

func TestGetPlayer(t *testing.T) {
	game := NewGame()
	player1 := game.AddPlayer("Spongebob")
	player2 := game.AddPlayer("Patrick")

	player3 := game.GetPlayerByName("Spongebob")

	if player1 != player3 {
		t.Error("getplayerbyName failed")
	}

	player4 := game.GetPlayerById(player2.Number)
	if player2 != player4 {
		t.Error("getplayerbyid failed")
	}

}

func contains(arr []Symbol, i Symbol) bool {
	for _, j := range arr {
		if j == i {
			return true
		}
	}
	return false
}
