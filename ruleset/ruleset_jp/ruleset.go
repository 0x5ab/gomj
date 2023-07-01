package ruleset_jp

import "github.com/0x5ab/gomj/ruleset"

var (
	JapaneseMahjongRuleset = &ruleset.Ruleset{
		YiZhongRuleset: YakuShuRuleset,
		PointRuleset:   FuRuleset,
	}
)
