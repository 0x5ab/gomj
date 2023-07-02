package gameplay

import (
	"github.com/0x5ab/gomj/wind"
)

type Game struct {
	Wind           wind.Wind
	Round          int
	DrawsRemaining int
}

func (g *Game) IsLastDraw() bool {
	// todo
	return g.DrawsRemaining == 0
}
