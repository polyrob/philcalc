package cards

import (
	"math/rand/v2"
)

type Deck struct {
	cards []Card
}

func NewDeck() Deck {
	deck := make([]Card, 52)
	i := 0
	for _, suit := range cardSuits {
		for _, cardValue := range cardValues {
			card := Card{value: cardValue, suit: suit}
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

func NewDeckWithoutDelt(alreadyDelt []Card) Deck {
	deck := make([]Card, 52-len(alreadyDelt))
	i := 0
	for _, suit := range cardSuits {
		for _, cardValue := range cardValues {
			card := Card{value: cardValue, suit: suit}
			delt := false
			for _, deltCard := range alreadyDelt {
				if card == deltCard {
					delt = true
					break
				}
			}
			if !delt {
				deck[i] = card
				i++
			}
		}
	}
	// shuffle deck
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return Deck{deck}
}

func (d *Deck) DealCard() Card {
	n := len(d.cards)
	poppedCard := (d.cards)[n-1]
	d.cards = (d.cards)[:n-1]
	return poppedCard
}
