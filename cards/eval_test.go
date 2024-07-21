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

	// 7s over 3s
	fullHouseMult3 = []card{
		{2, Spades},
		{3, Hearts},
		{3, Clubs},
		{7, Clubs},
		{7, Diamonds},
		{3, Clubs},
		{7, Hearts},
	}

	flush = []card{
		{7, Spades},
		{Ace, Hearts},
		{Queen, Hearts},
		{King, Hearts},
		{10, Hearts},
		{2, Clubs},
		{4, Hearts},
	}

	lowStraight = []card{
		{7, Spades},
		{Ace, Hearts},
		{4, Clubs},
		{3, Diamonds},
		{10, Hearts},
		{2, Clubs},
		{5, Hearts},
	}

	highStraight = []card{
		{7, Spades},
		{Ace, Hearts},
		{Queen, Diamonds},
		{King, Hearts},
		{10, Clubs},
		{2, Clubs},
		{Jack, Hearts},
	}

	threeOfAKind = []card{
		{Ace, Spades},
		{3, Hearts},
		{7, Clubs},
		{9, Clubs},
		{9, Diamonds},
		{9, Clubs},
		{2, Hearts},
	}

	twoPair = []card{
		{9, Diamonds},
		{King, Spades},
		{3, Hearts},
		{7, Clubs},
		{King, Clubs},
		{9, Clubs},
		{2, Hearts},
	}

	singlePair = []card{
		{9, Diamonds},
		{King, Spades},
		{3, Hearts},
		{Ace, Clubs},
		{King, Clubs},
		{8, Clubs},
		{2, Hearts},
	}

	highCard = []card{
		{Queen, Clubs},
		{10, Clubs},
		{9, Diamonds},
		{8, Clubs},
		{4, Spades},
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

func TestFullHouseMulti(t *testing.T) {
	eval := GetPokerHand(fullHouseMult3)
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

func TestFlush(t *testing.T) {
	eval := GetPokerHand(flush)
	assert.Equal(t, pokerHandType(Flush), eval.pokerHandType)

	expectedCards := []card{
		{Ace, Hearts},
		{King, Hearts},
		{Queen, Hearts},
		{10, Hearts},
		{4, Hearts},
	}
	assert.Equal(t, expectedCards, eval.cards)
}

func TestLowStraight(t *testing.T) {
	eval := GetPokerHand(lowStraight)
	assert.Equal(t, pokerHandType(Straight), eval.pokerHandType)

	expectedCards := []card{
		{5, Hearts},
		{4, Clubs},
		{3, Diamonds},
		{2, Clubs},
		{1, Hearts}, // Todo
	}
	assert.Equal(t, expectedCards, eval.cards)
}

func TestHighStraight(t *testing.T) {
	eval := GetPokerHand(highStraight)
	assert.Equal(t, pokerHandType(Straight), eval.pokerHandType)

	expectedCards := []card{
		{Ace, Hearts},
		{King, Hearts},
		{Queen, Diamonds},
		{Jack, Hearts},
		{10, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
}

func TestThreeOfAKind(t *testing.T) {
	eval := GetPokerHand(threeOfAKind)
	assert.Equal(t, pokerHandType(ThreeOfAKind), eval.pokerHandType)

	expectedCards := []card{
		{9, Clubs},
		{9, Diamonds},
		{9, Clubs},
		{Ace, Spades},
		{7, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
}

func TestTwoPair(t *testing.T) {
	eval := GetPokerHand(twoPair)
	assert.Equal(t, pokerHandType(TwoPair), eval.pokerHandType)

	expectedCards := []card{
		{King, Spades},
		{King, Clubs},
		{9, Diamonds},
		{9, Clubs},
		{7, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
}

func TestSinglePair(t *testing.T) {
	eval := GetPokerHand(singlePair)
	assert.Equal(t, pokerHandType(Pair), eval.pokerHandType)

	expectedCards := []card{
		{King, Spades},
		{King, Clubs},
		{Ace, Clubs},
		{9, Diamonds},
		{8, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
}

func TestHighCard(t *testing.T) {
	eval := GetPokerHand(highCard)
	assert.Equal(t, pokerHandType(HighCard), eval.pokerHandType)

	expectedCards := []card{
		{Queen, Clubs},
		{10, Clubs},
		{9, Diamonds},
		{8, Clubs},
		{4, Spades},
	}
	assert.Equal(t, expectedCards, eval.cards)
}
