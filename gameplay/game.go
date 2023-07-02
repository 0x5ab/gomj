package gameplay

import (
	"github.com/0x5ab/gomj/wind"
)

type Game struct {
	Wind           wind.Wind
	Round          int
	DrawsRemaining int
}

// IsLastTurn returns true if the current turn is the last turn of the game (last draw and last discard).
func (g *Game) IsLastTurn() bool {
	return g.DrawsRemaining == 0
}
