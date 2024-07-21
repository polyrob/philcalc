package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDealCards(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, 52, len(deck.cards))

	deltCard := deck.dealCard()
	assert.NotEqual(t, deltCard.suit, "")
	assert.Equal(t, len(deck.cards), 51)
}
