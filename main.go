package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"

	"philcalc/cards"
	"philcalc/runner"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	fmt.Printf("Starting")

	runs := flag.Int("runs", 100, "number of runs")
	playerInput := flag.String("player", "AhTd", "player hand")
	oppInput := flag.String("opp", "Jc7s", "opposing hand")
	flag.Parse()

	playerHand, err := parseHand(*playerInput)
	if err != nil {
		log.Fatal(err)
	}
	oppHand, err := parseHand(*oppInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Player hand: %s, opp hand: %s\n", playerHand, oppHand)

	s := runner.RunSim(*runs, *playerHand, *oppHand)
	s.PrintStats()
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
