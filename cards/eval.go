package cards

import (
	"sort"
)

const (
	HighCard = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type pokerHandType int

type Eval struct {
	pokerHandType pokerHandType
	cards         []card // best 5 cards
}

func GetPokerHand(cards []card) Eval {
	// sort by highest value
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].value > cards[j].value
	})

	// check if strightflush
	isFlush, flushCards := evalFlush(cards)
	if isFlush {
		// check if straight flush using just those cards
		isStraight, straighFlushCards := evalStraight(flushCards)
		if isStraight {
			// use only 5 highest
			playingCards := getTopCards(straighFlushCards)
			return Eval{
				pokerHandType: StraightFlush,
				cards:         playingCards,
			}
		}
	}

	// check occurrences of pairing cards
	cardSets := evalPairings(cards)

	// check 4 of a kind
	for _, cardSet := range cardSets {
		if len(cardSet) == 4 {
			// find the kicker
			for _, c := range cards {
				if c != cardSet[0] {
					// this is the next highest card not in of-a-kind set
					cardSet = append(cardSet, c)
					return Eval{
						pokerHandType: FourOfAKind,
						cards:         cardSet,
					}
				}
			}
		}
	}

	// check full house
	for _, cardSet := range cardSets {
		if len(cardSet) == 3 {
			// find another pair
			for _, subCardSet := range cardSets {
				if subCardSet[0].value == cardSet[0].value {
					continue // ignore the 3 of a kind already
				}

				if len(subCardSet) == 2 {
					cardSet = append(cardSet, subCardSet...)
					return Eval{
						pokerHandType: FullHouse,
						cards:         cardSet,
					}
				}
			}
		}
	}

	// check flush
	if isFlush {
		playingCards := getTopCards(flushCards)
		return Eval{
			pokerHandType: Flush,
			cards:         playingCards,
		}
	}

	// check straight
	isStraight, straightCards := evalStraight(cards)
	if isStraight {
		// use only 5 highest
		playingCards := getTopCards(straightCards)
		return Eval{
			pokerHandType: StraightFlush,
			cards:         playingCards,
		}
	}

	return Eval{}
}

func evalPairings(cards []card) map[int][]card {
	cardSets := make(map[int][]card)
	for _, c := range cards {
		v := c.value
		if set, ok := cardSets[v]; ok {
			cardSets[v] = append(set, c)
		} else {
			cardSets[v] = []card{c}
		}
	}
	return cardSets
}

// already in order presumably
func getTopCards(cards []card) []card {
	return cards[0:5]
}

func evalStraight(cards []card) (bool, []card) {
	if len(cards) < 5 {
		return false, nil
	}

	for _, c := range cards {
		if c.value == 14 {
			cards = append(cards, card{value: 1}) // add a "1" card with no suit
			break
		}
	}

	var valuesOnly []int
	for _, c := range cards {
		valuesOnly = append(valuesOnly, c.value)
	}
	//sort.Sort(sort.Reverse(sort.IntSlice(valuesOnly)))
	currentValue := valuesOnly[0]
	inARow := 1
	for i := 1; i < len(valuesOnly); i++ {
		if valuesOnly[i] == currentValue {
			continue
		}
		if valuesOnly[i] == currentValue-1 {
			inARow++
			if inARow == 5 {
				return true, cards
			}
		} else {
			inARow = 1
		}
		currentValue = valuesOnly[i]
	}

	return false, nil
}

func evalFlush(cards []card) (bool, []card) {
	if len(cards) < 5 {
		return false, nil
	}

	var (
		hearts   []card
		diamonds []card
		clubs    []card
		spades   []card
	)

	for _, card := range cards {
		switch card.suit {
		case "h":
			hearts = append(hearts, card)
			break
		case "d":
			diamonds = append(diamonds, card)
			break
		case "c":
			clubs = append(clubs, card)
			break
		case "s":
			spades = append(spades, card)
			break
		}
	}

	if len(hearts) >= 5 {
		return true, hearts
	}
	if len(diamonds) >= 5 {
		return true, diamonds
	}
	if len(clubs) >= 5 {
		return true, clubs
	}
	if len(spades) >= 5 {
		return true, spades
	}

	return false, nil
}
