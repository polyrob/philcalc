package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	straightFlush = []Card{
		{7, Spades},
		{Ace, Hearts},
		{Queen, Hearts},
		{King, Hearts},
		{10, Hearts},
		{2, Clubs},
		{Jack, Hearts},
	}

	fourOfAKind = []Card{
		{7, Spades},
		{Ace, Hearts},
		{Queen, Spades},
		{7, Clubs},
		{7, Diamonds},
		{2, Clubs},
		{7, Hearts},
	}

	fullHouse = []Card{
		{2, Spades},
		{3, Hearts},
		{3, Clubs},
		{7, Clubs},
		{7, Diamonds},
		{2, Clubs},
		{7, Hearts},
	}

	// 7s over 3s
	fullHouseMult3 = []Card{
		{2, Spades},
		{3, Hearts},
		{3, Clubs},
		{7, Clubs},
		{7, Diamonds},
		{3, Clubs},
		{7, Hearts},
	}

	flush = []Card{
		{7, Spades},
		{Ace, Hearts},
		{Queen, Hearts},
		{King, Hearts},
		{10, Hearts},
		{2, Clubs},
		{4, Hearts},
	}

	lowStraight = []Card{
		{7, Spades},
		{Ace, Hearts},
		{4, Clubs},
		{3, Diamonds},
		{10, Hearts},
		{2, Clubs},
		{5, Hearts},
	}

	highStraight = []Card{
		{7, Spades},
		{Ace, Hearts},
		{Queen, Diamonds},
		{King, Hearts},
		{10, Clubs},
		{2, Clubs},
		{Jack, Hearts},
	}

	threeOfAKind = []Card{
		{Ace, Spades},
		{3, Hearts},
		{7, Clubs},
		{9, Clubs},
		{9, Diamonds},
		{9, Clubs},
		{2, Hearts},
	}

	twoPair = []Card{
		{9, Diamonds},
		{King, Spades},
		{3, Hearts},
		{7, Clubs},
		{King, Clubs},
		{9, Clubs},
		{2, Hearts},
	}

	singlePair = []Card{
		{9, Diamonds},
		{King, Spades},
		{3, Hearts},
		{Ace, Clubs},
		{King, Clubs},
		{8, Clubs},
		{2, Hearts},
	}

	highCardQueen = []Card{
		{Queen, Clubs},
		{10, Clubs},
		{9, Diamonds},
		{8, Clubs},
		{4, Spades},
	}

	highCardQueenJack = []Card{
		{Queen, Clubs},
		{2, Clubs},
		{9, Diamonds},
		{Jack, Clubs},
		{4, Spades},
	}
)

func TestStraightFlush(t *testing.T) {
	eval := GetPokerHand(straightFlush)
	assert.Equal(t, pokerHandType(StraightFlush), eval.pokerHandType)

	expectedCards := []Card{
		{Ace, Hearts},
		{King, Hearts},
		{Queen, Hearts},
		{Jack, Hearts},
		{10, Hearts},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Straight Flush - Ace high", eval.String())
}

func TestFourOfAKind(t *testing.T) {
	eval := GetPokerHand(fourOfAKind)
	assert.Equal(t, pokerHandType(FourOfAKind), eval.pokerHandType)

	expectedCards := []Card{
		{7, Spades},
		{7, Clubs},
		{7, Diamonds},
		{7, Hearts},
		{Ace, Hearts},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Four-Of-A-Kind - 7s", eval.String())
}

func TestFullHouse(t *testing.T) {
	eval := GetPokerHand(fullHouse)
	assert.Equal(t, pokerHandType(FullHouse), eval.pokerHandType)

	expectedCards := []Card{
		{7, Clubs},
		{7, Diamonds},
		{7, Hearts},
		{3, Hearts},
		{3, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Full House - 7s full of 3s", eval.String())
}

func TestFullHouseMulti(t *testing.T) {
	eval := GetPokerHand(fullHouseMult3)
	assert.Equal(t, pokerHandType(FullHouse), eval.pokerHandType)

	expectedCards := []Card{
		{7, Clubs},
		{7, Diamonds},
		{7, Hearts},
		{3, Hearts},
		{3, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Full House - 7s full of 3s", eval.String())
}

func TestFlush(t *testing.T) {
	eval := GetPokerHand(flush)
	assert.Equal(t, pokerHandType(Flush), eval.pokerHandType)

	expectedCards := []Card{
		{Ace, Hearts},
		{King, Hearts},
		{Queen, Hearts},
		{10, Hearts},
		{4, Hearts},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Flush (Hearts) - Ace high", eval.String())
}

func TestLowStraight(t *testing.T) {
	eval := GetPokerHand(lowStraight)
	assert.Equal(t, pokerHandType(Straight), eval.pokerHandType)

	expectedCards := []Card{
		{5, Hearts},
		{4, Clubs},
		{3, Diamonds},
		{2, Clubs},
		{1, Hearts}, // Todo okay keeping the ace is a different Card struct?
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Straight - 5 high", eval.String())
}

func TestHighStraight(t *testing.T) {
	eval := GetPokerHand(highStraight)
	assert.Equal(t, pokerHandType(Straight), eval.pokerHandType)

	expectedCards := []Card{
		{Ace, Hearts},
		{King, Hearts},
		{Queen, Diamonds},
		{Jack, Hearts},
		{10, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Straight - Ace high", eval.String())
}

func TestThreeOfAKind(t *testing.T) {
	eval := GetPokerHand(threeOfAKind)
	assert.Equal(t, pokerHandType(ThreeOfAKind), eval.pokerHandType)

	expectedCards := []Card{
		{9, Clubs},
		{9, Diamonds},
		{9, Clubs},
		{Ace, Spades},
		{7, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Three-Of-A-Kind - 9s", eval.String())
}

func TestTwoPair(t *testing.T) {
	eval := GetPokerHand(twoPair)
	assert.Equal(t, pokerHandType(TwoPair), eval.pokerHandType)

	expectedCards := []Card{
		{King, Spades},
		{King, Clubs},
		{9, Diamonds},
		{9, Clubs},
		{7, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Two Pair - Kings and 9s with a 7 kicker", eval.String())
}

func TestSinglePair(t *testing.T) {
	eval := GetPokerHand(singlePair)
	assert.Equal(t, pokerHandType(Pair), eval.pokerHandType)

	expectedCards := []Card{
		{King, Spades},
		{King, Clubs},
		{Ace, Clubs},
		{9, Diamonds},
		{8, Clubs},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "Pair - Kings with a Ace kicker", eval.String())
}

func TestHighCard(t *testing.T) {
	eval := GetPokerHand(highCardQueen)
	assert.Equal(t, pokerHandType(HighCard), eval.pokerHandType)

	expectedCards := []Card{
		{Queen, Clubs},
		{10, Clubs},
		{9, Diamonds},
		{8, Clubs},
		{4, Spades},
	}
	assert.Equal(t, expectedCards, eval.cards)
	assert.Equal(t, "High Card - Queen high", eval.String())
}

func TestBeats(t *testing.T) {
	highCardQueenE := GetPokerHand(highCardQueen)
	singlePairE := GetPokerHand(singlePair)
	twoPairE := GetPokerHand(twoPair)
	threeOfAKindE := GetPokerHand(threeOfAKind)
	lowStraightE := GetPokerHand(lowStraight)
	highStraightE := GetPokerHand(highStraight)
	flushE := GetPokerHand(flush)
	fullHouseE := GetPokerHand(fullHouse)
	fourOfAKindE := GetPokerHand(fourOfAKind)
	straightFlushE := GetPokerHand(straightFlush)

	assert.Positive(t, Beats(straightFlushE, fourOfAKindE))
	assert.Positive(t, Beats(fourOfAKindE, fullHouseE))
	assert.Positive(t, Beats(fullHouseE, flushE))
	assert.Positive(t, Beats(flushE, highStraightE))
	assert.Positive(t, Beats(highStraightE, lowStraightE))
	assert.Positive(t, Beats(lowStraightE, threeOfAKindE))
	assert.Positive(t, Beats(threeOfAKindE, twoPairE))
	assert.Positive(t, Beats(twoPairE, singlePairE))
	assert.Positive(t, Beats(singlePairE, highCardQueenE))
	assert.Positive(t, Beats(GetPokerHand(highCardQueenJack), highCardQueenE))

	assert.Negative(t, Beats(highCardQueenE, singlePairE))

	assert.Zero(t, Beats(highCardQueenE, highCardQueenE))
}
