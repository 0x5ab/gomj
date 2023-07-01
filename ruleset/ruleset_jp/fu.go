package ruleset_jp

import "github.com/0x5ab/gomj/ruleset"

type JapaneseMahjongFuRuleset struct{}

var (
	FuRuleset = &JapaneseMahjongFuRuleset{}
)

func (r *JapaneseMahjongFuRuleset) GetPoint(huWay *ruleset.HuWay) int {
	return 10
}
