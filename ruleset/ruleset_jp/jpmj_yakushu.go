package ruleset_jp

import (
	"github.com/0x5ab/gomj/gameplay"
	"github.com/0x5ab/gomj/ruleset"
)

const (
	// 1番
	IdRiiChi      = "jp_riichi"
	IdIppatsu     = "jp_ippatsu"
	IdMenQianZimo = "jp_mqz"
	IdTanYao      = "jp_tanyao"
	IdPingHu      = "jp_ph"
	IdYiBeiKou    = "jp_ybk"

	// 2番
	IdSanSeTongShun = "jp_ssts"
	IdDuiDuiHu      = "jp_ddh"
	IdQingYiSe      = "jp_qys"

	// 役满
	IdShiSanYao = "jp_ssy"
	IdZiYiSe    = "jp_zys"
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

func (j *JapaneseMahjongYiZhongRuleset) Check(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) []ruleset.YiZhong {
	var yiZhongs []ruleset.YiZhong
	isMenQing := hand.IsMenQing()
	for _, y := range j.YiZhongs {
		if !isMenQing && y.NeedMenQing() {
			continue
		}
		if y.IsYiZhong(huWay, hand, gotTile) {
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
			&RiiChi{},
			&Ippatsu{},
			&TanYao{},
			&MenQianZimo{},
			&ToiToi{},
			&PinHu{},
			&QingYiSe{},
			&ShiSanYao{},
			&ZiYiSe{},
		},
	}
)

// #region riichi

type RiiChi struct{}

func (r *RiiChi) GetId() string {
	return IdRiiChi
}

func (r *RiiChi) GetFan(hand *gameplay.Hand) int {
	return 1
}

func (r *RiiChi) IsYakuman() bool {
	return false
}

func (r *RiiChi) GetName() string {
	return "立直"
}

func (r *RiiChi) NeedMenQing() bool {
	return true
}

func (r *RiiChi) GetDescription() string {
	return "立直"
}

func (r *RiiChi) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return hand.IsRiichi
}

// #endregion

// #region ippatsu

type Ippatsu struct{}

func (i *Ippatsu) GetId() string {
	return IdIppatsu
}

func (i *Ippatsu) GetFan(hand *gameplay.Hand) int {
	return 1
}

func (i *Ippatsu) IsYakuman() bool {
	return false
}

func (i *Ippatsu) GetName() string {
	return "一发"
}

func (i *Ippatsu) NeedMenQing() bool {
	return true
}

func (i *Ippatsu) GetDescription() string {
	return "玩家立直后，自己摸入的第一只牌即自摸胡，或者在这之间食胡他人打出的牌。但中途遇上其他玩家鸣牌则无效。又称“即”，部分竞技麻雀不采用此规则。"
}

func (i *Ippatsu) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return hand.DrawsAfterRiichi == 1
}

// #endregion

// #region menqianzimo

type MenQianZimo struct{}

func (m *MenQianZimo) GetId() string {
	return IdMenQianZimo
}

func (m *MenQianZimo) GetFan(hand *gameplay.Hand) int {
	return 1
}

func (m *MenQianZimo) IsYakuman() bool {
	return false
}

func (m *MenQianZimo) GetName() string {
	return "门前清自摸和"
}

func (m *MenQianZimo) NeedMenQing() bool {
	return true
}

func (m *MenQianZimo) GetDescription() string {
	return "门清听牌的状态下自摸和牌。"
}

func (m *MenQianZimo) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return hand.IsMenQing() && gotTile.State() == gameplay.GameTileDrawn
}

// #endregion

// #region tanyao

type TanYao struct{}

func (t *TanYao) GetId() string {
	return IdTanYao
}

func (t *TanYao) GetFan(hand *gameplay.Hand) int {
	return 1
}

func (t *TanYao) IsYakuman() bool {
	return false
}

func (t *TanYao) GetName() string {
	return "断幺九"
}

func (t *TanYao) NeedMenQing() bool {
	return false
}

func (t *TanYao) GetDescription() string {
	return "由非幺九牌组成的和牌"
}

func (t *TanYao) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return hand.IsTanYao()
}

// #endregion

// #region pinhu

type PinHu struct{}

func (p *PinHu) GetId() string {
	return IdPingHu
}

func (p *PinHu) GetFan(hand *gameplay.Hand) int {
	if hand.IsMenQing() {
		return 2
	}
	return 1
}

func (p *PinHu) IsYakuman() bool {
	return false
}

func (p *PinHu) GetName() string {
	return "平和"
}

func (p *PinHu) NeedMenQing() bool {
	return true
}

func (p *PinHu) GetDescription() string {
	return "门清状态下由4副顺子和将牌组成的和牌"
}

func (p *PinHu) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return huWay.GetShunZiCount() == 4
}

// #endregion

// #region toitoi

type ToiToi struct{}

func (d *ToiToi) GetId() string {
	return IdDuiDuiHu
}

func (d *ToiToi) GetFan(hand *gameplay.Hand) int {
	return 2
}

func (d *ToiToi) IsYakuman() bool {
	return false
}

func (d *ToiToi) GetName() string {
	return "对对胡"
}

func (d *ToiToi) NeedMenQing() bool {
	return false
}

func (d *ToiToi) GetDescription() string {
	return "由4副刻子（杠）和将牌组成的和牌"
}

func (d *ToiToi) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return huWay.GetKeZiCount()+hand.GetGangCount()+hand.GetPengCount() == 4
}

// #endregion

// #region sansetongshun

type SanSeTongShun struct{}

func (s *SanSeTongShun) GetId() string {
	return IdSanSeTongShun
}

func (s *SanSeTongShun) GetFan(hand *gameplay.Hand) int {
	return 2
}

func (s *SanSeTongShun) IsYakuman() bool {
	return false
}

func (s *SanSeTongShun) GetName() string {
	return "三色同顺"
}

func (s *SanSeTongShun) NeedMenQing() bool {
	return false
}

func (s *SanSeTongShun) GetDescription() string {
	return "三色同顺"
}

// #endregion

// #region qingyise

type QingYiSe struct{}

func (q *QingYiSe) GetId() string {
	return IdQingYiSe
}

func (q *QingYiSe) GetFan(hand *gameplay.Hand) int {
	if hand.IsMenQing() {
		return 6
	}
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

func (q *QingYiSe) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return hand.IsQingYiSe()
}

// #endregion

// #region shisanyao

type ShiSanYao struct{}

func (s *ShiSanYao) GetId() string {
	return IdShiSanYao
}

func (s *ShiSanYao) GetFan(hand *gameplay.Hand) int {
	return 13
}

func (s *ShiSanYao) IsYakuman() bool {
	return true
}

func (s *ShiSanYao) GetName() string {
	return "国士无双"
}

func (s *ShiSanYao) NeedMenQing() bool {
	return true
}

func (s *ShiSanYao) GetDescription() string {
	return "全数为单只幺九牌，第14只则可为其中一只幺九牌。"
}

func (s *ShiSanYao) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return huWay.IsShiSanYao && !huWay.IsShiSanYaoDanJi
}

// #endregion

// #region ziyise

type ZiYiSe struct{}

func (z *ZiYiSe) GetId() string {
	return IdZiYiSe
}

func (z *ZiYiSe) GetFan(hand *gameplay.Hand) int {
	return 13
}

func (z *ZiYiSe) IsYakuman() bool {
	return true
}

func (z *ZiYiSe) GetName() string {
	return "字一色"
}

func (z *ZiYiSe) NeedMenQing() bool {
	return false
}

func (z *ZiYiSe) GetDescription() string {
	return "由字牌组成的和牌"
}

func (z *ZiYiSe) IsYiZhong(huWay *ruleset.HuWay, hand *gameplay.Hand, gotTile *gameplay.GameTile) bool {
	return hand.IsZiYiSe()
}

// #endregion
