package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandDisplay(t *testing.T) {
	h := Hand{
		Cards: []Card{
			{4, "h"},
			{10, "d"},
			{11, "c"},
			{12, "s"},
			{14, "h"},
		},
	}
	display := h.String()
	assert.Equal(t, "[4 of Hearts], [10 of Diamonds], [Jack of Clubs], [Queen of Spades], [Ace of Hearts]", display)
}
