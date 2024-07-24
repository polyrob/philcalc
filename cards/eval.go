package cards

import (
	"fmt"
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

func (w pokerHandType) String() string {
	return [...]string{"High Card", "Pair", "Two Pair", "Three-Of-A-Kind", "Straight", "Flush", "Full House", "Four-Of-A-Kind", "Straight Flush"}[w]
}

type pokerHandType int

type Eval struct {
	pokerHandType pokerHandType
	cards         []Card // best 5 cards
}

func GetPokerHand(cards []Card) Eval {
	// sort by highest value
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].value > cards[j].value
	})

	// check if straight flush
	isFlush, flushCards := evalFlush(cards)
	if isFlush {
		// check if straight flush using just those Cards
		isStraight, straightFlushCards := evalStraight(flushCards)
		if isStraight {
			// use only 5 highest
			playingCards := getTopCards(straightFlushCards)
			return Eval{
				pokerHandType: StraightFlush,
				cards:         playingCards,
			}
		}
	}

	// check occurrences of pairing Cards
	cardSets := evalPairings(cards)

	// check 4 of a kind
	for _, cardSet := range cardSets {
		if len(cardSet) == 4 {
			// find the kicker
			for _, c := range cards {
				if c != cardSet[0] {
					// this is the next highest Card not in of-a-kind set
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
			var bestPair []Card
			for _, subCardSet := range cardSets {
				if subCardSet[0].value == cardSet[0].value {
					continue // ignore the 3 of a kind already
				}

				if len(subCardSet) >= 2 {
					if bestPair == nil || subCardSet[0].value > bestPair[0].value {
						bestPair = subCardSet[0:2]
					}
				}
			}

			if bestPair != nil {
				cardSet = append(cardSet, bestPair...)
				return Eval{
					pokerHandType: FullHouse,
					cards:         cardSet,
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
			pokerHandType: Straight,
			cards:         playingCards,
		}
	}

	// check 3 of a kind
	for _, cardSet := range cardSets {
		if len(cardSet) == 3 {
			// best kickers
			for _, c := range cards {
				if c.value != cardSet[0].value {
					cardSet = append(cardSet, c)
					if len(cardSet) == 5 {
						return Eval{
							pokerHandType: ThreeOfAKind,
							cards:         cardSet,
						}
					}
				}
			}
		}
	}

	// check two pair
	var bestCards []Card
	var pairs [][]Card
	for _, cardSet := range cardSets {
		if len(cardSet) == 2 {
			pairs = append(pairs, cardSet)
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0].value > pairs[j][0].value
	})

	if len(pairs) > 1 {
		// find another pair
		bestCards = append(bestCards, pairs[0]...)
		bestCards = append(bestCards, pairs[1]...)
		// best kicker
		for _, c := range cards {
			if c.value != bestCards[0].value && c.value != bestCards[2].value {
				bestCards = append(bestCards, c)
				return Eval{
					pokerHandType: TwoPair,
					cards:         bestCards,
				}
			}
		}
	} else if len(pairs) == 1 {
		// just a pair then
		bestCards = append(bestCards, pairs[0]...)
		for _, c := range cards {
			if c.value != bestCards[0].value {
				bestCards = append(bestCards, c)
				if len(bestCards) == 5 {
					return Eval{
						pokerHandType: Pair,
						cards:         bestCards,
					}
				}
			}
		}
	}

	// high Card
	return Eval{
		HighCard,
		cards[:5],
	}
}

func evalPairings(cards []Card) map[int][]Card {
	cardSets := make(map[int][]Card)
	for _, c := range cards {
		v := c.value
		if set, ok := cardSets[v]; ok {
			cardSets[v] = append(set, c)
		} else {
			cardSets[v] = []Card{c}
		}
	}
	return cardSets
}

// already in order presumably
func getTopCards(cards []Card) []Card {
	return cards[0:5]
}

func evalStraight(cards []Card) (bool, []Card) {
	if len(cards) < 5 {
		return false, nil
	}

	for _, c := range cards {
		if c.value == 14 {
			cards = append(cards, Card{value: 1, suit: c.suit}) // fake a "1" Card
		}
	}

	currentValue := cards[0].value
	var inARow = []Card{cards[0]}
	for i := 1; i < len(cards); i++ {
		if cards[i].value == currentValue {
			continue
		}
		if cards[i].value == currentValue-1 {
			inARow = append(inARow, cards[i])
			if len(inARow) == 5 {
				return true, inARow
			}
		} else {
			inARow = []Card{cards[i]}
		}
		currentValue = cards[i].value
	}

	return false, nil
}

func evalFlush(cards []Card) (bool, []Card) {
	if len(cards) < 5 {
		return false, nil
	}

	var (
		hearts   []Card
		diamonds []Card
		clubs    []Card
		spades   []Card
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

func (e Eval) String() string {
	switch e.pokerHandType {
	case FullHouse:
		return fmt.Sprintf("%s - %ss full of %ss", e.pokerHandType.String(), e.cards[0].DisplayValue(), e.cards[4].DisplayValue())
	case FourOfAKind:
		return fmt.Sprintf("%s - %ss", e.pokerHandType.String(), e.cards[0].DisplayValue())
	case ThreeOfAKind:
		return fmt.Sprintf("%s - %ss", e.pokerHandType.String(), e.cards[0].DisplayValue())
	case TwoPair:
		return fmt.Sprintf("%s - %ss and %ss with a %s kicker", e.pokerHandType.String(), e.cards[0].DisplayValue(), e.cards[2].DisplayValue(), e.cards[4].DisplayValue())
	case Flush:
		return fmt.Sprintf("%s (%s) - %s high", e.pokerHandType.String(), e.cards[0].DisplaySuit(), e.cards[0].DisplayValue())
	case Pair:
		return fmt.Sprintf("%s - %ss with a %s kicker", e.pokerHandType.String(), e.cards[0].DisplayValue(), e.cards[2].DisplayValue())
	}

	return fmt.Sprintf("%s - %s high", e.pokerHandType.String(), e.cards[0].DisplayValue())
}

func Beats(e1, e2 Eval) int {
	if e1.pokerHandType != e2.pokerHandType {
		if e1.pokerHandType > e2.pokerHandType {
			return 1
		} else {
			return -1
		}
	}

	// if Cards are of same type, look for high Card
	for i := 0; i < len(e1.cards); i++ {
		if e1.cards[i].value == e2.cards[i].value {
			continue
		}
		if e1.cards[i].value > e2.cards[i].value {
			return 1
		} else {
			return -1
		}
	}
	return 0 // push
}
