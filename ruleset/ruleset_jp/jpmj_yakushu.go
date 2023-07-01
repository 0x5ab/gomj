package ruleset_jp

import (
	"github.com/0x5ab/gomj/gameplay"
	"github.com/0x5ab/gomj/ruleset"
)

const (
	IdDuiDuiHu        = "jp_ddh"
	IdPingHu          = "jp_ph"
	IdQingYiSe        = "jp_qys"
	IdQingYiSeMenQing = "jp_qysmq"
)

func ContainsYiZhong(yiZhongs []ruleset.YiZhong, yiZhongId string) bool {
	for _, y := range yiZhongs {
		if y.GetId() == IdDuiDuiHu {
			return true
		}
	}
	return false
}

type JapaneseMahjongYiZhong interface {
	ruleset.YiZhong
	IsYakuman() bool
	NeedMenQing() bool
}

type JapaneseMahjongYiZhongRuleset struct {
	YiZhongs []JapaneseMahjongYiZhong
}

func (j *JapaneseMahjongYiZhongRuleset) GetAllYiZhongs() []ruleset.YiZhong {
	yiZhongs := make([]ruleset.YiZhong, len(j.YiZhongs))
	for i, y := range j.YiZhongs {
		yiZhongs[i] = y
	}
	return yiZhongs
}

func (j *JapaneseMahjongYiZhongRuleset) Check(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.PlayedTile) []ruleset.YiZhong {
	var yiZhongs []ruleset.YiZhong
	isMenQing := hand.IsMenQing()
	for _, y := range j.YiZhongs {
		if !isMenQing && y.NeedMenQing() {
			continue
		}
		if y.GetId() == IdQingYiSeMenQing {
			// did this in qingyise
			continue
		}
		if y.IsYiZhong(huWay, hand, gotTile) {
			if y.GetId() == IdQingYiSe {
				// check if qingyisemenqing
				if isMenQing {
					yiZhongs = append(yiZhongs, &QingYiSeMenQing{})
					continue
				}
			}
			yiZhongs = append(yiZhongs, y)
		}
	}
	if ContainsYiZhong(yiZhongs, IdDuiDuiHu) {
		// check for duplicates
	}
	return yiZhongs
}

var (
	YiZhongRuleset = JapaneseMahjongYiZhongRuleset{
		YiZhongs: []JapaneseMahjongYiZhong{
			&DuiDuiHu{},
			&PingHu{},
			&QingYiSe{},
			&QingYiSeMenQing{},
		},
	}
)

type DuiDuiHu struct{}

func (d *DuiDuiHu) GetId() string {
	return IdDuiDuiHu
}

func (d *DuiDuiHu) GetFan() int {
	return 2
}

func (d *DuiDuiHu) IsYakuman() bool {
	return false
}

func (d *DuiDuiHu) GetName() string {
	return "对对胡"
}

func (d *DuiDuiHu) NeedMenQing() bool {
	return false
}

func (d *DuiDuiHu) GetDescription() string {
	return "由4副刻子（杠）和将牌组成的和牌"
}

func (d *DuiDuiHu) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.PlayedTile) bool {
	return huWay.GetKeZiCount()+hand.GetGangCount()+hand.GetPengCount() == 4
}

type PingHu struct{}

func (p *PingHu) GetId() string {
	return IdPingHu
}

func (p *PingHu) GetFan() int {
	return 1
}

func (p *PingHu) IsYakuman() bool {
	return false
}

func (p *PingHu) GetName() string {
	return "平和"
}

func (p *PingHu) NeedMenQing() bool {
	return true
}

func (p *PingHu) GetDescription() string {
	return "门清状态下由4副顺子和将牌组成的和牌"
}

func (p *PingHu) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.PlayedTile) bool {
	return huWay.GetShunZiCount() == 4
}

type QingYiSeMenQing struct{}

func (q *QingYiSeMenQing) GetId() string {
	return IdQingYiSeMenQing
}

func (q *QingYiSeMenQing) GetFan() int {
	return 6
}

func (q *QingYiSeMenQing) IsYakuman() bool {
	return false
}

func (q *QingYiSeMenQing) GetName() string {
	return "清一色"
}

func (q *QingYiSeMenQing) NeedMenQing() bool {
	return true
}

func (q *QingYiSeMenQing) GetDescription() string {
	return "由一种花色的牌组成的和牌"
}

func (q *QingYiSeMenQing) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.PlayedTile) bool {
	return hand.IsMenQing() && hand.IsQingYiSe()
}

type QingYiSe struct{}

func (q *QingYiSe) GetId() string {
	return IdQingYiSe
}

func (q *QingYiSe) GetFan() int {
	return 5
}

func (q *QingYiSe) IsYakuman() bool {
	return false
}

func (q *QingYiSe) GetName() string {
	return "清一色"
}

func (q *QingYiSe) NeedMenQing() bool {
	return false
}

func (q *QingYiSe) GetDescription() string {
	return "由一种花色的牌组成的和牌"
}

func (q *QingYiSe) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.PlayedTile) bool {
	return hand.IsQingYiSe()
}
