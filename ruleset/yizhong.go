package ruleset

import (
	"github.com/0x5ab/gomj/gameplay"
)

type YiZhongRuleset interface {
	GetAllYiZhongs() []YiZhong
	Check(huWay *HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) []YiZhong
}

type YiZhong interface {
	GetId() string
	GetFan(hand *gameplay.Hand) int
	GetName() string
	GetDescription() string
	IsYiZhong(huWay *HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool
}
