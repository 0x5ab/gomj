package gameplay

import (
	"github.com/0x5ab/gomj/wind"
)

type Game struct {
	Id               int
	Wind             wind.Wind
	Round            int
	DrawsRemaining   int
	HasMingPai       bool // 是否有玩家鸣牌（吃、碰、杠、或其它特殊操作（拔北等））
	SpecialGameProps any
}

// IsLastTurn returns true if the current turn is the last turn of the game (last draw and last discard).
func (g *Game) IsLastTurn() bool {
	return g.DrawsRemaining == 0
}
