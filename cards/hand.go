package cards

import (
	"fmt"
	"strings"
)

type Hand struct {
	cards []card
}

func (h *Hand) acceptCard(c card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) String() string {
	var cardStrings []string
	for _, c := range h.cards {
		cardStrings = append(cardStrings, c.String())
	}
	return fmt.Sprintf(strings.Join(cardStrings, ", "))
}
