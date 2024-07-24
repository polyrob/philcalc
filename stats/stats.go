package stats

import (
	"fmt"
	"time"

	"philcalc/cards"
)

type Stats struct {
	Start, End           time.Time
	RoundsPlayed         int
	Wins, Losses, Pushes int

	BestWinningHand *cards.Eval
	BestLosingHand  *cards.Eval
	BestLostTo      *cards.Eval
}

func (s *Stats) GetExecutionTime() time.Duration {
	return s.End.Sub(s.Start)
}

func (s *Stats) PrintStats() {
	fmt.Printf("\tWins: %d,\tPct: %.2f\n", s.Wins, float32(s.Wins)/float32(s.RoundsPlayed)*100)
	fmt.Printf("\tLosses: %d,\tPct: %.2f\n", s.Losses, float32(s.Losses)/float32(s.RoundsPlayed)*100)
	fmt.Printf("\tPushes: %d,\tPct: %.2f\n", s.Pushes, float32(s.Pushes)/float32(s.RoundsPlayed)*100)

	if s.BestWinningHand != nil {
		fmt.Printf("\tBest Winning Hand: %s\n", s.BestWinningHand)
	}
	if s.BestLosingHand != nil {
		fmt.Printf("\tBest Losing Hand: %s\n", s.BestLosingHand)
		fmt.Printf("\t\tLost to: %s\n", s.BestLostTo)
	}
}
