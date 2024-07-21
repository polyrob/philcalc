package main

import (
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Starting")
	//args := os.Args
	//
	//d := newDeck()
	//slog.Debug("got deck", d)
	//
	//playerHand := hand{}
	//dealCard(&d, &playerHand)
	//dealCard(&d, &playerHand)
	//
	//// deal community
	//communityCards := hand{}
	//dealCard(&d, &communityCards)
	//dealCard(&d, &communityCards)
	//dealCard(&d, &communityCards)
	//
	////strait test
	//dealCard(&d, &communityCards)
	//dealCard(&d, &communityCards)
	//
	//// turn
	//dealCard(&d, &communityCards)
	//
	//// river
	//dealCard(&d, &communityCards)
	//
	//slog.Debug(playerHand.getBest(communityCards))
	//
	//slog.Debug("got cards", playerHand)
	slog.Info("Complete")
}
