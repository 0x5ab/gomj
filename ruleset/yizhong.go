package ruleset

import (
	"github.com/0x5ab/gomj/gameplay"
)

type YiZhongRuleset interface {
	GetAllYiZhongs() []YiZhong
	Check(huWay *HuWay, hand *gameplay.Hand, gotTile *gameplay.PlayedTile) []YiZhong
}

type YiZhong interface {
	GetId() string
	GetFan() int
	GetName() string
	GetDescription() string
	IsYiZhong(huWay *HuWay, hand *gameplay.Hand, gotTile *gameplay.PlayedTile) bool
}
