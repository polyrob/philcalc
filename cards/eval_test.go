package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	straightFlush = []card{
		{7, Spades},
		{Ace, Hearts},
		{Queen, Hearts},
		{King, Hearts},
		{10, Hearts},
		{2, Clubs},
		{Jack, Hearts},
	}

	fourOfAKind = []card{
		{7, Spades},
		{Ace, Hearts},
		{Queen, Spades},
		{7, Clubs},
		{7, Diamonds},
		{2, Clubs},
		{7, Hearts},
	}

	fullHouse = []card{
		{2, Spades},
		{3, Hearts},
		{3, Clubs},
		{7, Clubs},
		{7, Diamonds},
		{2, Clubs},
		{7, Hearts},
	}
)

func TestStraightFlush(t *testing.T) {
	eval := GetPokerHand(straightFlush)
	assert.Equal(t, pokerHandType(StraightFlush), eval.pokerHandType)

	expectedCards := []card{
		{Ace, Hearts},
		{King, Hearts},
		{Queen, Hearts},
		{Jack, Hearts},
		{10, Hearts},
	}
	assert.Equal(t, expectedCards, eval.cards)
}

func TestFourOfAKind(t *testing.T) {
	eval := GetPokerHand(fourOfAKind)
	assert.Equal(t, pokerHandType(FourOfAKind), eval.pokerHandType)

	expectedCards := []card{
		{7, Spades},
		{7, Clubs},
		{7, Diamonds},
		{7, Hearts},
		{Ace, Hearts},
	}
	assert.Equal(t, expectedCards, eval.cards)
}

func TestFullHouse(t *testing.T) {
	eval := GetPokerHand(fullHouse)
	assert.Equal(t, pokerHandType(FullHouse), eval.pokerHandType)

	expectedCards := []card{
		{7, Clubs},
		{7, Diamonds},
		{7, Hearts},
		{3, Hearts},
		{3, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
}
