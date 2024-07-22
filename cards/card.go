package cards

import (
	"errors"
	"fmt"
	"strconv"
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
	Jack:  "J",
	Queen: "Q",
	King:  "K",
	Ace:   "A",
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

type Card struct {
	value int
	suit  string
}

func (c Card) DisplayValue() string {
	return cardValueMap[c.value]
}

func (c Card) DisplaySuit() string {
	return cardSuitMap[c.suit]
}

func (c Card) String() string {
	return fmt.Sprintf("[%s of %s]", cardValueMap[c.value], cardSuitMap[c.suit])
}

func ParseCard(s string) (*Card, error) {
	if len(s) != 2 {
		return nil, errors.New("invalid card")
	}
	valStr := string(s[0])
	suit := string(s[1])

	val, err := strconv.Atoi(valStr)
	if err != nil {
		if valStr == "A" {
			val = Ace
		} else if valStr == "K" {
			val = King
		} else if valStr == "Q" {
			val = Queen
		} else if valStr == "J" {
			val = Jack
		} else if valStr == "T" {
			val = 10
		}
	}
	return &Card{
		value: val,
		suit:  suit,
	}, nil
}
