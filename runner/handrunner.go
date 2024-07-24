package runner

import (
	"fmt"
	"time"

	"philcalc/cards"
	"philcalc/stats"
)

func RunSim(iterations int, playerHand, otherHnad cards.Hand) stats.Stats {
	simStats := stats.Stats{}

	simStats.Start = time.Now()
	for i := 0; i < iterations; i++ {
		simStats.RoundsPlayed++

		deck := cards.NewDeckWithoutDelt(append(playerHand.Cards, otherHnad.Cards...))

		community := cards.Hand{}
		// flop
		community.AcceptCard(deck.DealCard())
		community.AcceptCard(deck.DealCard())
		community.AcceptCard(deck.DealCard())

		// turn
		community.AcceptCard(deck.DealCard())

		// river
		community.AcceptCard(deck.DealCard())

		pEval := cards.GetPokerHand(append(playerHand.Cards, community.Cards...))
		oEval := cards.GetPokerHand(append(otherHnad.Cards, community.Cards...))
		result := cards.Beats(pEval, oEval)
		if result > 0 {
			simStats.Wins++
			if simStats.BestWinningHand == nil || cards.Beats(pEval, *simStats.BestWinningHand) > 0 {
				simStats.BestWinningHand = &pEval
				fmt.Printf("\nNew Best Winning Hand, %s\n", pEval.String())
			}
		} else if result < 0 {
			simStats.Losses++
			if simStats.BestLosingHand == nil || cards.Beats(pEval, *simStats.BestLosingHand) > 0 {
				simStats.BestLosingHand = &pEval
				simStats.BestLostTo = &oEval
				fmt.Printf("\nNew Best Losing Hand, %s\n", pEval.String())
			}
		} else {
			simStats.Pushes++
		}

		fmt.Printf(".")
		if (i+1)%100 == 0 {
			fmt.Printf("\n")
		}
	}
	simStats.End = time.Now()

	fmt.Printf("Sim Complete in %s\n", simStats.GetExecutionTime())
	return simStats
}
