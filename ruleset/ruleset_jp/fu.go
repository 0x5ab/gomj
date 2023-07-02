package ruleset_jp

import (
	"math"

	"github.com/0x5ab/gomj/ruleset"
)

type JapaneseMahjongFuRuleset struct{}

var (
	FuRuleset = &JapaneseMahjongFuRuleset{}
)

func (r *JapaneseMahjongFuRuleset) GetPoint(huWay *ruleset.HuWay) int {
	if huWay.IsQiDui {
		return 25
	}
	fu := 20
	// 听牌
	if huWay.GetTingPaiType() == ruleset.TingPaiTypeSpecial {
		fu += 2
	}
	// 雀头
	if huWay.QueTou.IsWindType(huWay.Hand.Game.Wind) ||
		huWay.QueTou.IsWindType(huWay.Hand.Player.Wind) {
		fu += 2
		if huWay.Hand.Game.Wind == huWay.Hand.Player.Wind {
			// 连风雀头，有的规则是4符
			fu += 2
		}
	} else if huWay.QueTou.IsSanYuan() {
		// 三元牌
		fu += 2
	}
	// 刻子
	// 暗刻
	for _, kezi := range huWay.Kezi {
		if kezi.IsYaoJiu() {
			fu += 8
		} else {
			fu += 4
		}
	}
	// 明刻
	for _, kezi := range huWay.Hand.GetKeZiTiles() {
		if kezi.IsYaoJiu() {
			fu += 4
		} else {
			fu += 2
		}
	}
	// 暗杠
	for _, gang := range huWay.Hand.GetAnGangTiles() {
		if gang.IsYaoJiu() {
			fu += 32
		} else {
			fu += 16
		}
	}
	// 明杠
	for _, gang := range huWay.Hand.GetMingGangTiles() {
		if gang.IsYaoJiu() {
			fu += 16
		} else {
			fu += 8
		}
	}
	return int(math.Round(math.Ceil(float64(fu)/10) * 10))
}
