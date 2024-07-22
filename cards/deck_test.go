package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDealCards(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, 52, len(deck.cards))

	deltCard := deck.DealCard()
	assert.NotEqual(t, deltCard.suit, "")
	assert.Equal(t, len(deck.cards), 51)
}

func TestNewDeckWithoutDelt(t *testing.T) {
	alreadyDeltCards := []Card{
		{Ace, Hearts},
		{5, Diamonds},
	}

	deck := NewDeckWithoutDelt(alreadyDeltCards)
	assert.Equal(t, 50, len(deck.cards))
	for _, c := range deck.cards {
		for _, delt := range alreadyDeltCards {
			if c == delt {
				assert.Fail(t, "This Card should have already been delt and not in the deck", delt)
			}
		}
	}
}
