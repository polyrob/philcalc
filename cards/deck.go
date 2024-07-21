package cards

import (
	"math/rand/v2"
)

type Deck struct {
	cards []card
}

func NewDeck() Deck {
	deck := make([]card, 52)
	i := 0
	for _, suit := range cardSuits {
		for _, cardValue := range cardValues {
			card := card{value: cardValue, suit: suit}
			deck[i] = card
			i++
		}
	}
	// shuffle deck
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return Deck{deck}
}

func (d *Deck) dealCard() card {
	n := len(d.cards)
	poppedCard := (d.cards)[n-1]
	d.cards = (d.cards)[:n-1]
	return poppedCard
}
