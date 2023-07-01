package gameplay

type Wind int

const (
	East  Wind = 1
	South Wind = 2
	West  Wind = 3
	North Wind = 4
)

type Game struct {
	Wind           Wind
	Round          int
	DrawsRemaining int
}

func (g *Game) IsLastDraw() bool {
	// todo
	return g.DrawsRemaining == 0
}
