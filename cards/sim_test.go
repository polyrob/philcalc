package cards

import (
	"fmt"
	"testing"
)

func TestSimSingle(t *testing.T) {
	myCards := []card{
		{Ace, Hearts},
		{5, Diamonds},
	}

	deck := NewDeckWithoutDelt(myCards)

	var otherPlayersCards []card
	// deal other play random cards
	otherPlayersCards = append(otherPlayersCards, deck.dealCard())
	otherPlayersCards = append(otherPlayersCards, deck.dealCard())

	var communityCards []card
	// flow
	communityCards = append(communityCards, deck.dealCard())
	communityCards = append(communityCards, deck.dealCard())
	communityCards = append(communityCards, deck.dealCard())

	// turn
	communityCards = append(communityCards, deck.dealCard())

	// river
	communityCards = append(communityCards, deck.dealCard())

	myEval := GetPokerHand(append(myCards, communityCards...))
	otherPlayerEval := GetPokerHand(append(otherPlayersCards, communityCards...))

	fmt.Printf("You cards are %s\n", myCards)
	fmt.Printf("Other players cards are %s\n", otherPlayersCards)
	fmt.Printf("The community cards are %s\n", communityCards)

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

type WinsAndLosses struct {
	wins   int
	losses int
	pushes int
}

func (r WinsAndLosses) GetDetails() string {
	total := r.wins + r.losses + r.pushes
	return fmt.Sprintf("%.2f win pct, %.2f loss pct, %.2f tie pct", float32(r.wins)/float32(total)*100, float32(r.losses)/float32(total)*100, float32(r.pushes)/float32(total)*100)
}

func TestSimPermutations(t *testing.T) {
	runsPerHand := 20000
	suited := Hearts
	nonSuited := Spades
	statMap := make(map[string]WinsAndLosses)

	for i, firstCard := range cardValues {
		card1 := card{
			value: firstCard,
			suit:  suited,
		}

		// suited
		for j := i + 1; j < len(cardValues); j++ {
			card2 := card{
				value: cardValues[j],
				suit:  suited,
			}
			myCards := []card{card1, card2}

			key := fmt.Sprintf("%s-%s suited", card2.DisplayValue(), card1.DisplayValue())
			fmt.Printf("Running %s\n", key)
			for runCount := 0; runCount < runsPerHand; runCount++ {
				result := getSimResult(myCards)
				stats, ok := statMap[key]
				if !ok {
					stats = WinsAndLosses{}
				}
				if result.outcome > 0 {
					stats.wins++
				} else if result.outcome < 0 {
					stats.losses++
				} else {
					stats.pushes++
				}
				statMap[key] = stats
			}
		}

		// unsuited
		for j := i; j < len(cardValues); j++ {
			card2 := card{
				value: cardValues[j],
				suit:  nonSuited,
			}
			myCards := []card{card1, card2}

			key := fmt.Sprintf("%s-%s unsuited", card2.DisplayValue(), card1.DisplayValue())
			fmt.Printf("Running %s\n", key)
			for runCount := 0; runCount < runsPerHand; runCount++ {
				result := getSimResult(myCards)
				stats, ok := statMap[key]
				if !ok {
					stats = WinsAndLosses{}
				}
				if result.outcome > 0 {
					stats.wins++
				} else if result.outcome < 0 {
					stats.losses++
				} else {
					stats.pushes++
				}
				statMap[key] = stats
			}
		}
	}

	fmt.Println("Complete.")
	for k, v := range statMap {
		fmt.Printf("%s: %s\n", k, v.GetDetails())
	}
	fmt.Println("stop here.")
}

func getSimResult(playerCards []card) SimResult {
	deck := NewDeckWithoutDelt(playerCards)

	var otherPlayersCards []card
	// deal other play random cards
	otherPlayersCards = append(otherPlayersCards, deck.dealCard())
	otherPlayersCards = append(otherPlayersCards, deck.dealCard())

	var communityCards []card
	// flow
	communityCards = append(communityCards, deck.dealCard())
	communityCards = append(communityCards, deck.dealCard())
	communityCards = append(communityCards, deck.dealCard())

	// turn
	communityCards = append(communityCards, deck.dealCard())

	// river
	communityCards = append(communityCards, deck.dealCard())

	myEval := GetPokerHand(append(playerCards, communityCards...))
	otherPlayerEval := GetPokerHand(append(otherPlayersCards, communityCards...))

	//fmt.Printf("You cards are %s\n", playerCards)
	//fmt.Printf("Other players cards are %s\n", otherPlayersCards)
	//fmt.Printf("The community cards are %s\n", communityCards)

	outcome := Beats(myEval, otherPlayerEval)
	return SimResult{
		outcome:  outcome,
		handType: myEval.pokerHandType,
	}
	//if outcome > 0 {
	//	return SimResult{
	//		outcome: outcome,
	//	}
	//} else if outcome < 0 {
	//	fmt.Println("You lost.")
	//} else {
	//	fmt.Println("You pushed.")
	//}
}

type SimResult struct {
	outcome  int
	handType pokerHandType
}
