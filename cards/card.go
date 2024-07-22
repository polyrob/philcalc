package cards

import (
	"fmt"
)

var (
	cardValues = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	cardSuits  = []string{Hearts, Diamonds, Clubs, Spades}
)

var cardSuitMap = map[string]string{
	Hearts:   "Hearts",
	Diamonds: "Diamonds",
	Clubs:    "Clubs",
	Spades:   "Spades",
}

var cardValueMap = map[int]string{
	1:     "Ace",
	2:     "2",
	3:     "3",
	4:     "4",
	5:     "5",
	6:     "6",
	7:     "7",
	8:     "8",
	9:     "9",
	10:    "10",
	Jack:  "Jack",
	Queen: "Queen",
	King:  "King",
	Ace:   "Ace",
}

const (
	Ace      = 14
	King     = 13
	Queen    = 12
	Jack     = 11
	Hearts   = "h"
	Diamonds = "d"
	Clubs    = "c"
	Spades   = "s"
)

type card struct {
	value int
	suit  string
}

func (c card) DisplayValue() string {
	return cardValueMap[c.value]
}

func (c card) DisplaySuit() string {
	return cardSuitMap[c.suit]
}

func (c card) String() string {
	return fmt.Sprintf("[%s of %s]", cardValueMap[c.value], cardSuitMap[c.suit])
}
