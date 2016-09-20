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

func contains(arr []Symbol, i Symbol) bool {
	for _, j := range arr {
		if j == i {
			return true
		}
	}
	return false
}
