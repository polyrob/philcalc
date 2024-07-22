package cards

import (
	"fmt"
	"testing"
)

type WinsAndLosses struct {
	wins   int
	losses int
	pushes int
}

func (r WinsAndLosses) GetDetails() string {
	total := r.wins + r.losses + r.pushes
	return fmt.Sprintf("%.2f win pct, %.2f loss pct, %.2f tie pct", float32(r.wins)/float32(total)*100, float32(r.losses)/float32(total)*100, float32(r.pushes)/float32(total)*100)
}

func TestSimSingle(t *testing.T) {
	myCards := []Card{
		{Ace, Hearts},
		{5, Diamonds},
	}

	deck := NewDeckWithoutDelt(myCards)

	var otherPlayersCards []Card
	// deal other play random Cards
	otherPlayersCards = append(otherPlayersCards, deck.DealCard())
	otherPlayersCards = append(otherPlayersCards, deck.DealCard())

	var communityCards []Card
	// flow
	communityCards = append(communityCards, deck.DealCard())
	communityCards = append(communityCards, deck.DealCard())
	communityCards = append(communityCards, deck.DealCard())

	// turn
	communityCards = append(communityCards, deck.DealCard())

	// river
	communityCards = append(communityCards, deck.DealCard())

	myEval := GetPokerHand(append(myCards, communityCards...))
	otherPlayerEval := GetPokerHand(append(otherPlayersCards, communityCards...))

	fmt.Printf("You Cards are %s\n", myCards)
	fmt.Printf("Other players Cards are %s\n", otherPlayersCards)
	fmt.Printf("The community Cards are %s\n", communityCards)

	outcome := Beats(myEval, otherPlayerEval)
	if outcome > 0 {
		fmt.Println("You win!")
	} else if outcome < 0 {
		fmt.Println("You lost.")
	} else {
		fmt.Println("You pushed.")
	}

	fmt.Printf("\tyour hand is: %s.\n\tother player hand is: %s\n", myEval.String(), otherPlayerEval.String())
}

func getSimResult(playerCards []Card) SimResult {
	deck := NewDeckWithoutDelt(playerCards)

	var otherPlayersCards []Card
	// deal other play random Cards
	otherPlayersCards = append(otherPlayersCards, deck.DealCard())
	otherPlayersCards = append(otherPlayersCards, deck.DealCard())

	var communityCards []Card
	// flow
	communityCards = append(communityCards, deck.DealCard())
	communityCards = append(communityCards, deck.DealCard())
	communityCards = append(communityCards, deck.DealCard())

	// turn
	communityCards = append(communityCards, deck.DealCard())

	// river
	communityCards = append(communityCards, deck.DealCard())

	myEval := GetPokerHand(append(playerCards, communityCards...))
	otherPlayerEval := GetPokerHand(append(otherPlayersCards, communityCards...))

	outcome := Beats(myEval, otherPlayerEval)
	return SimResult{
		outcome:  outcome,
		handType: myEval.pokerHandType,
	}
}

type SimResult struct {
	outcome  int
	handType pokerHandType
}
