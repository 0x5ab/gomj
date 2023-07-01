package ruleset

import (
	"github.com/0x5ab/gomj/gameplay"
)

type YiZhong interface {
	GetId() string
	GetFan(hand *gameplay.Hand) int
	GetName() string
	GetDescription() string
	IsYiZhong(huWay *HuWay) bool
}

type YiZhongRuleset interface {
	GetAllYiZhongs() []YiZhong
	Check(huWay *HuWay) []YiZhong
}

type PointRuleset interface {
	GetPoint(huWay *HuWay) int
}

type Ruleset struct {
	YiZhongRuleset // nonnull
	PointRuleset   // nullable
}

func (r *Ruleset) GetYiZhongs(huWay *HuWay) []YiZhong {
	return r.YiZhongRuleset.Check(huWay)
}

func (r *Ruleset) GetPoint(huWay *HuWay) int {
	if r.PointRuleset != nil {
		return r.PointRuleset.GetPoint(huWay)
	}
	return 0
}
