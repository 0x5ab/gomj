package ruleset_jp

import (
	"github.com/0x5ab/gomj/gameplay"
	"github.com/0x5ab/gomj/ruleset"
	"github.com/0x5ab/gomj/tiles"
)

const (
	// Special
	IdDora    = "jp_dora"
	IdUraDora = "jp_uradora"
	IdAkaDora = "jp_akadora"
	IdPei     = "jp_pei"

	// 1番
	IdRiiChi      = "jp_riichi"
	IdIppatsu     = "jp_ippatsu"
	IdMenQianZimo = "jp_mqz"
	IdTanYao      = "jp_tanyao"
	IdPingHu      = "jp_ph"
	IdYiBeiKou    = "jp_ybk"
	IdYakuHai     = "jp_yh"
	IdLingShang   = "jp_ls"
	IdQiangGang   = "jp_qg"
	IdHaiDi       = "jp_haidi"
	IdHeDi        = "jp_hedi"

	// 2番
	IdSanSeTongShun = "jp_ssts"
	IdSanSeTongKe   = "jp_sstk"
	IdYiTiaoLong    = "jp_ytl"
	IdDuiDuiHu      = "jp_ddh"
	IdSanAnKe       = "jp_sanak"
	IdSanGangZi     = "jp_sangz"
	IdQiDuiZi       = "jp_qdz"
	IdHunQuanDaiYao = "jp_hqdy"
	IdHunLaoTou     = "jp_hlt"
	IdXiaoSanYuan   = "jp_xsy"
	IdDoubleRiiChi  = "jp_drc"

	// 3番
	IdHunYiSe        = "jp_hys"
	IdChunQuanDaiYao = "jp_cqdy"
	IdErBeiKou       = "jp_ebk"

	// 6番
	IdQingYiSe = "jp_qys"

	// 役满
	IdShiSanYao      = "jp_ssy"
	IdShiSanYaoDanJi = "jp_ssydj"
	IdDaSanYuan      = "jp_dsy"
	IdSiAnKe         = "jp_siak"
	IdSiAnKeDanJi    = "jp_siakdj"
	IdZiYiSe         = "jp_zys"
	IdLvYiSe         = "jp_lys"
	IdXiaoSiXi       = "jp_xsx"
	IdDaSiXi         = "jp_dsx"
	IdQingLaoTou     = "jp_qlt"
	IdJiuLianBaoDeng = "jp_jlbd"
	IdChunJiuLian    = "jp_cjl"
	IdSiGangZi       = "jp_sigz"
	IdTianHu         = "jp_th"
	IdDiHu           = "jp_dh"
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

func (j *JapaneseMahjongYiZhongRuleset) Check(huWay *ruleset.HuWay) []ruleset.YiZhong {
	var yiZhongs []ruleset.YiZhong
	isMenQing := huWay.Hand.IsMenQing()
	for _, y := range j.YiZhongs {
		if !isMenQing && y.NeedMenQing() {
			continue
		}
		if y.IsYiZhong(huWay) {
			yiZhongs = append(yiZhongs, y)
		}
	}
	if ContainsYiZhong(yiZhongs, IdDuiDuiHu) {
		// check for duplicates
	}
	return yiZhongs
}

var (
	YakuShuRuleset = &JapaneseMahjongYiZhongRuleset{
		YiZhongs: []JapaneseMahjongYiZhong{
			&RiiChi{},
			&Ippatsu{},
			&TanYao{},
			&MenQianZimo{},
			&ToiToi{},
			&SanSeTongShun{},
			&SanSeTongKe{},
			&PinHu{},
			&YiBeiKou{},
			&QingYiSe{},
			&ShiSanYao{},
			&ShiSanYaoDanJi{},
			&SiAnKe{},
			&SiAnKeDanJi{},
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

func (r *RiiChi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsRiichi()
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

func (i *Ippatsu) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.Hand.DrawsAfterRiichi == 1
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

func (m *MenQianZimo) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsMenQing() && huWay.GotTile.IsJustDrawn()
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

func (t *TanYao) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsTanYao()
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

func (p *PinHu) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.GetShunZiCountWithoutChi() == 4
}

// #endregion

// #region yibeikou

type YiBeiKou struct{}

func (y *YiBeiKou) GetId() string {
	return IdYiBeiKou
}

func (y *YiBeiKou) GetFan(hand *gameplay.Hand) int {
	return 1
}

func (y *YiBeiKou) IsYakuman() bool {
	return false
}

func (y *YiBeiKou) GetName() string {
	return "一杯口"
}

func (y *YiBeiKou) NeedMenQing() bool {
	return true
}

func (y *YiBeiKou) GetDescription() string {
	return "由两副相同的顺子组成的和牌"
}

func (y *YiBeiKou) IsYiZhong(huWay *ruleset.HuWay) bool {
	return tiles.CountDuplicateTiles(huWay.Shunzi) == 1
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
	return "由4副刻子（杠）和雀头组成的和牌"
}

func (d *ToiToi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.GetKeZiCount() == 4
}

// #endregion

// #region sansetongshun

type SanSeTongShun struct{}

func (s *SanSeTongShun) GetId() string {
	return IdSanSeTongShun
}

func (s *SanSeTongShun) GetFan(hand *gameplay.Hand) int {
	if hand.IsMenQing() {
		return 2
	}
	return 1
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

func (s *SanSeTongShun) IsYiZhong(huWay *ruleset.HuWay) bool {
	if huWay.GetShunZiCount() < 3 {
		return false
	}
	hasStartTiles := make(map[int]int)
	hasWan := false
	hasSuo := false
	hasTong := false
	for _, shunZi := range huWay.GetAllShunZi() {
		if shunZi.TileType == tiles.Wan {
			hasWan = true
		}
		if shunZi.TileType == tiles.Suo {
			hasSuo = true
		}
		if shunZi.TileType == tiles.Tong {
			hasTong = true
		}
		hasStartTiles[shunZi.Number]++
	}
	if !hasWan || !hasSuo || !hasTong {
		return false
	}
	for _, v := range hasStartTiles {
		if v == 3 {
			return true
		}
	}
	return false
}

// #endregion

// #region sansetongke

type SanSeTongKe struct{}

func (s *SanSeTongKe) GetId() string {
	return IdSanSeTongKe
}

func (s *SanSeTongKe) GetFan(hand *gameplay.Hand) int {
	return 2
}

func (s *SanSeTongKe) IsYakuman() bool {
	return false
}

func (s *SanSeTongKe) GetName() string {
	return "三色同刻"
}

func (s *SanSeTongKe) NeedMenQing() bool {
	return false
}

func (s *SanSeTongKe) GetDescription() string {
	return "三色同刻"
}

func (s *SanSeTongKe) IsYiZhong(huWay *ruleset.HuWay) bool {
	if huWay.GetKeZiCount() < 3 {
		return false
	}
	hasStartTiles := make(map[int]int)
	hasWan := false
	hasSuo := false
	hasTong := false
	for _, keZi := range huWay.GetAllKeZi() {
		if keZi.TileType == tiles.Wan {
			hasWan = true
		}
		if keZi.TileType == tiles.Suo {
			hasSuo = true
		}
		if keZi.TileType == tiles.Tong {
			hasTong = true
		}
		hasStartTiles[keZi.Number]++
	}
	if !hasWan || !hasSuo || !hasTong {
		return false
	}
	for _, v := range hasStartTiles {
		if v == 3 {
			return true
		}
	}
	return false
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

func (q *QingYiSe) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsQingYiSe()
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

func (s *ShiSanYao) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsShiSanYao && !huWay.IsShiSanYaoDanJi
}

// #endregion

// #region shisanyaodanji

type ShiSanYaoDanJi struct{}

func (s *ShiSanYaoDanJi) GetId() string {
	return IdShiSanYaoDanJi
}

func (s *ShiSanYaoDanJi) GetFan(hand *gameplay.Hand) int {
	return 13
}

func (s *ShiSanYaoDanJi) IsYakuman() bool {
	return true
}

func (s *ShiSanYaoDanJi) GetName() string {
	return "国士无双十三面"
}

func (s *ShiSanYaoDanJi) NeedMenQing() bool {
	return true
}

func (s *ShiSanYaoDanJi) GetDescription() string {
	return "全数为单只幺九牌，第14只则可为其中一只幺九牌。"
}

func (s *ShiSanYaoDanJi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsShiSanYao && huWay.IsShiSanYaoDanJi
}

// #endregion

// #region sianke

type SiAnKe struct{}

func (s *SiAnKe) GetId() string {
	return IdSiAnKe
}

func (s *SiAnKe) GetFan(hand *gameplay.Hand) int {
	return 13
}

func (s *SiAnKe) IsYakuman() bool {
	return true
}

func (s *SiAnKe) GetName() string {
	return "四暗刻"
}

func (s *SiAnKe) NeedMenQing() bool {
	return true
}

func (s *SiAnKe) GetDescription() string {
	return "四组暗刻"
}

func (s *SiAnKe) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.GetKeZiCountWithoutPeng() == 4 && huWay.IsGotTileQueTou()
}

// #endregion

// #region siankedanji

type SiAnKeDanJi struct{}

func (s *SiAnKeDanJi) GetId() string {
	return IdSiAnKe
}

func (s *SiAnKeDanJi) GetFan(hand *gameplay.Hand) int {
	return 13
}

func (s *SiAnKeDanJi) IsYakuman() bool {
	return true
}

func (s *SiAnKeDanJi) GetName() string {
	return "四暗刻单骑"
}

func (s *SiAnKeDanJi) NeedMenQing() bool {
	return true
}

func (s *SiAnKeDanJi) GetDescription() string {
	return "四组暗刻，且只胡一张牌（单钓）"
}

func (s *SiAnKeDanJi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.GetKeZiCount() == 4 && huWay.IsGotTileQueTou()
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

func (z *ZiYiSe) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.Hand.IsZiYiSe()
}

// #endregion
