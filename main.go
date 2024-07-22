package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"

	"philcalc/cards"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	fmt.Printf("Starting")

	runs := flag.Int("runs", 100, "number of runs")
	playerInput := flag.String("player", "AhTd", "player hand")
	oppInput := flag.String("opp", "Jc7s", "opposing hand")
	flag.Parse()

	fmt.Println("runs:", *runs)
	fmt.Println("player:", *playerInput)
	fmt.Println("opp:", *oppInput)

	playerHand, err := parseHand(*playerInput)
	if err != nil {
		log.Fatal(err)
	}
	oppHand, err := parseHand(*oppInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Player hand: %s, opp hand: %s\n", playerHand, oppHand)

	wins, losses, pushes := 0, 0, 0

	fmt.Printf("Staring %d runs\n", *runs)
	for i := 0; i < *runs; i++ {

		deck := cards.NewDeckWithoutDelt(append(playerHand.Cards, oppHand.Cards...))

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
		oEval := cards.GetPokerHand(append(oppHand.Cards, community.Cards...))
		result := cards.Beats(pEval, oEval)
		if result > 0 {
			wins++
		} else if result < 0 {
			losses++
		} else {
			pushes++
		}

		fmt.Printf(".")
		if (i+1)%100 == 0 {
			fmt.Printf("\n")
		}
	}

	fmt.Println("Complete")

	fmt.Printf("\tWins: %d,\tPct: %.2f\n", wins, float32(wins)/float32(*runs)*100)
	fmt.Printf("\tLosses: %d,\tPct: %.2f\n", losses, float32(losses)/float32(*runs)*100)
	fmt.Printf("\tPushes: %d,\tPct: %.2f\n", pushes, float32(pushes)/float32(*runs)*100)
}

func parseHand(s string) (*cards.Hand, error) {
	card1, err := cards.ParseCard(s[0:2])
	if err != nil {
		return nil, err
	}
	card2, err := cards.ParseCard(s[2:4])
	if err != nil {
		return nil, err
	}

	return &cards.Hand{
		[]cards.Card{
			*card1,
			*card2,
		},
	}, nil
}
