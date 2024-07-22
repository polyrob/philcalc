package cards

import (
	"fmt"
	"strings"
)

type Hand struct {
	Cards []Card
}

func (h *Hand) AcceptCard(c Card) {
	h.Cards = append(h.Cards, c)
}

func (h *Hand) String() string {
	var cardStrings []string
	for _, c := range h.Cards {
		cardStrings = append(cardStrings, c.String())
	}
	return fmt.Sprintf(strings.Join(cardStrings, ", "))
}
