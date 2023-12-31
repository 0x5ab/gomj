package gameplay

import "github.com/0x5ab/gomj/wind"

type Player struct {
	Game *Game
	Wind wind.Wind
	Id   int
	Hand Hand
}

func (p *Player) Equal(p2 *Player) bool {
	return p != nil && p2 != nil && p.Id == p2.Id
}

func (p *Player) IsZhuangJia() bool {
	return p.Wind == p.Game.Wind
}
