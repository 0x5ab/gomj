package gameplay

import "github.com/0x5ab/gomj/wind"

type Player struct {
	Game *Game
	Wind wind.Wind
	Id   int
	Hand Hand
}
