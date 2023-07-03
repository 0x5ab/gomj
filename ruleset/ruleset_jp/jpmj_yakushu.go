package ruleset_jp

import (
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
	IdToiToi        = "jp_tt"
	IdSanAnKe       = "jp_sanak"
	IdSanGangZi     = "jp_sangz"
	IdQiDuiZi       = "jp_qdz"
	IdHunQuanDaiYao = "jp_hqdy"
	IdHunLaoTou     = "jp_hlt"
	IdXiaoSanYuan   = "jp_xsy"
	IdDoubleRiichi  = "jp_drc"

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

func YakuEquals(y JapaneseMahjongYaku, yid string) bool {
	return y.GetId() == yid
}

func ContainsYiZhong(yiZhongs []JapaneseMahjongYaku, yakuId string) bool {
	for _, y := range yiZhongs {
		if YakuEquals(y, yakuId) {
			return true
		}
	}
	return false
}

func FilterYakuman(yiZhongs []JapaneseMahjongYaku) []JapaneseMahjongYaku {
	var filtered []JapaneseMahjongYaku
	for _, y := range yiZhongs {
		if y.YakumanBaisu() > 0 {
			filtered = append(filtered, y)
		}
	}
	return filtered
}

func ConvertToYiZhong(jYakus []JapaneseMahjongYaku) []ruleset.YiZhong {
	yiZhongs := make([]ruleset.YiZhong, len(jYakus))
	for i, y := range jYakus {
		yiZhongs[i] = y
	}
	return yiZhongs
}

func GetFuLuMinusOneFan(fan int, huWay *ruleset.HuWay) int {
	if huWay.IsMenQing() {
		return fan
	}
	return fan - 1
}

type JapaneseMahjongYaku interface {
	ruleset.YiZhong
	YakumanBaisu() int
	NeedMenQing() bool
}

type JapaneseMahjongYakuRuleset struct {
	Yakus []JapaneseMahjongYaku
}

func (j *JapaneseMahjongYakuRuleset) GetAllYiZhongs() []ruleset.YiZhong {
	yiZhongs := make([]ruleset.YiZhong, len(j.Yakus))
	for i, y := range j.Yakus {
		yiZhongs[i] = y
	}
	return yiZhongs
}

func (j *JapaneseMahjongYakuRuleset) Check(huWay *ruleset.HuWay) []ruleset.YiZhong {
	var yakus []JapaneseMahjongYaku
	isMenQing := huWay.Hand.IsMenQing()
	for _, y := range j.Yakus {
		if !isMenQing && y.NeedMenQing() {
			continue
		}
		if y.IsYiZhong(huWay) {
			yakus = append(yakus, y)
		}
	}
	yakumans := FilterYakuman(yakus)
	if len(yakumans) > 0 {
		return ConvertToYiZhong(yakumans)
	}

	return ConvertToYiZhong(yakus)
}

var (
	YakuShuRuleset = &JapaneseMahjongYakuRuleset{
		Yakus: []JapaneseMahjongYaku{
			&RiiChi{},
			&Ippatsu{},
			&TanYao{},
			&MenQianZimo{},
			&PinHu{},
			&YiBeiKou{},
			&YakuHai{},
			&LingShang{},
			&QiangGang{},
			&HaiDi{},
			&HeDi{},
			&SanSeTongShun{},
			&SanSeTongKe{},
			&YiTiaoLong{},
			&ToiToi{},
			&SanAnKe{},
			&SanGangZi{},
			&QiDuiZi{},
			&HunQuanDaiYao{},
			&HunLaoTou{},
			&XiaoSanYuan{},
			&DoubleRiichi{},
			&HunYiSe{},
			&ChunQuanDaiYao{},
			&ErBeiKou{},
			&QingYiSe{},
			&ShiSanYao{},
			&ShiSanYaoDanJi{},
			&DaSanYuan{},
			&SiAnKe{},
			&SiAnKeDanJi{},
			&ZiYiSe{},
			&LvYiSe{},
			&XiaoSiXi{},
			&DaSiXi{},
			&QingLaoTou{},
			&JiuLianBaoDeng{},
			&ChunJiuLian{},
			&SiGangZi{},
			&TianHu{},
			&DiHu{},
		},
	}
)

// #region riichi

type RiiChi struct{}

func (r *RiiChi) GetId() string {
	return IdRiiChi
}

func (r *RiiChi) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (r *RiiChi) YakumanBaisu() int {
	return 0
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
	return huWay.IsRiichi() && huWay.Hand.DrawNumber != 1
}

// #endregion

// #region ippatsu

type Ippatsu struct{}

func (i *Ippatsu) GetId() string {
	return IdIppatsu
}

func (i *Ippatsu) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (i *Ippatsu) YakumanBaisu() int {
	return 0
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

func (m *MenQianZimo) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (m *MenQianZimo) YakumanBaisu() int {
	return 0
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
	return huWay.IsMenQing() && huWay.IsZiMo()
}

// #endregion

// #region tanyao

type TanYao struct{}

func (t *TanYao) GetId() string {
	return IdTanYao
}

func (t *TanYao) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (t *TanYao) YakumanBaisu() int {
	return 0
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

func (p *PinHu) GetFan(huWay *ruleset.HuWay) int {
	return 1
}

func (p *PinHu) YakumanBaisu() int {
	return 0
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

func (y *YiBeiKou) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (y *YiBeiKou) YakumanBaisu() int {
	return 0
}

func (y *YiBeiKou) GetName() string {
	return "一盃口"
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

// #region yakuhai

type YakuHai struct{}

func (y *YakuHai) GetId() string {
	return IdYakuHai
}

func (y *YakuHai) GetFan(huWay *ruleset.HuWay) int {
	fan := 0
	for _, kezi := range huWay.GetAllKeZi() {
		if kezi.IsSanYuan() || huWay.IsTileZiFeng(&kezi) || huWay.IsTileChangFeng(&kezi) {
			fan++
		}
	}
	return fan
}

func (y *YakuHai) YakumanBaisu() int {
	return 0
}

func (y *YakuHai) GetName() string {
	return "役牌"
}

func (y *YakuHai) NeedMenQing() bool {
	return false
}

func (y *YakuHai) GetDescription() string {
	return "包括由三元牌、自风牌、场风牌组成的刻子或杠子。"
}

func (y *YakuHai) IsYiZhong(huWay *ruleset.HuWay) bool {
	for _, kezi := range huWay.GetAllKeZi() {
		if kezi.IsSanYuan() || huWay.IsTileZiFeng(&kezi) || huWay.IsTileChangFeng(&kezi) {
			return true
		}
	}
	return false
}

// #endregion

// #region lingshang

type LingShang struct{}

func (l *LingShang) GetId() string {
	return IdLingShang
}

func (l *LingShang) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (l *LingShang) YakumanBaisu() int {
	return 0
}

func (l *LingShang) GetName() string {
	return "岭上开花"
}

func (l *LingShang) NeedMenQing() bool {
	return false
}

func (l *LingShang) GetDescription() string {
	return "开杠后摸的岭上牌自摸和牌。"
}

func (l *LingShang) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsZiMo() && huWay.GotTile.IsLingShang
}

// #endregion

// #region qianggang

type QiangGang struct{}

func (q *QiangGang) GetId() string {
	return IdQiangGang
}

func (q *QiangGang) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (q *QiangGang) YakumanBaisu() int {
	return 0
}

func (q *QiangGang) GetName() string {
	return "抢杠"
}

func (q *QiangGang) NeedMenQing() bool {
	return false
}

func (q *QiangGang) GetDescription() string {
	return "荣和其他人加杠的牌。"
}

func (q *QiangGang) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.GotTile.IsLingShang && !huWay.IsZiMo()
}

// #endregion

// #region haidi

type HaiDi struct{}

func (h *HaiDi) GetId() string {
	return IdHaiDi
}

func (h *HaiDi) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (h *HaiDi) YakumanBaisu() int {
	return 0
}

func (h *HaiDi) GetName() string {
	return "海底捞月"
}

func (h *HaiDi) NeedMenQing() bool {
	return false
}

func (h *HaiDi) GetDescription() string {
	return "玩家摸到牌山最后一张牌而自摸胡。"
}

func (h *HaiDi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsZiMo() && huWay.Hand.Game.IsLastTurn()
}

// #endregion

// #region hedi

type HeDi struct{}

func (h *HeDi) GetId() string {
	return IdHeDi
}

func (h *HeDi) GetFan(_ *ruleset.HuWay) int {
	return 1
}

func (h *HeDi) YakumanBaisu() int {
	return 0
}

func (h *HeDi) GetName() string {
	return "河底捞鱼"
}

func (h *HeDi) NeedMenQing() bool {
	return false
}

func (h *HeDi) GetDescription() string {
	return "玩家荣和牌河中最后一张打出的牌。"
}

func (h *HeDi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return !huWay.IsZiMo() && huWay.Hand.Game.IsLastTurn()
}

// #endregion

// #region toitoi

type ToiToi struct{}

func (d *ToiToi) GetId() string {
	return IdToiToi
}

func (d *ToiToi) GetFan(_ *ruleset.HuWay) int {
	return 2
}

func (d *ToiToi) YakumanBaisu() int {
	return 0
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

func (s *SanSeTongShun) GetFan(huWay *ruleset.HuWay) int {
	return GetFuLuMinusOneFan(2, huWay)
}

func (s *SanSeTongShun) YakumanBaisu() int {
	return 0
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
	// TODO: needs to be optimized
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

func (s *SanSeTongKe) GetFan(_ *ruleset.HuWay) int {
	return 2
}

func (s *SanSeTongKe) YakumanBaisu() int {
	return 0
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

// #region yitiaolong

type YiTiaoLong struct{}

func (y *YiTiaoLong) GetId() string {
	return IdYiTiaoLong
}

func (y *YiTiaoLong) GetFan(huWay *ruleset.HuWay) int {
	return GetFuLuMinusOneFan(2, huWay)
}

func (y *YiTiaoLong) YakumanBaisu() int {
	return 0
}

func (y *YiTiaoLong) GetName() string {
	return "一气通贯"
}

func (y *YiTiaoLong) NeedMenQing() bool {
	return false
}

func (y *YiTiaoLong) GetDescription() string {
	return "同一色牌中（筒索万），一至九各有一只，组成三副顺子"
}

func (y *YiTiaoLong) IsYiZhong(huWay *ruleset.HuWay) bool {
	if huWay.GetShunZiCount() < 3 {
		return false
	}
	colors := [3]int{}
	color := -1
	for _, shunZi := range huWay.GetAllShunZi() {
		colors[shunZi.TileType.Index()]++
		if colors[shunZi.TileType.Index()] == 3 {
			color = shunZi.TileType.Index()
			break
		}
	}
	if color == -1 {
		return false
	}
	tileType := tiles.TileTypeFromIndex(color)
	hasOne := false
	hasFour := false
	hasSeven := false
	for _, shunZi := range huWay.GetAllShunZiWithTileType(tileType) {
		switch shunZi.Number {
		case 1:
			hasOne = true
		case 4:
			hasFour = true
		case 7:
			hasSeven = true
		}
	}
	return hasOne && hasFour && hasSeven
}

// #endregion

// #region sananke

type SanAnKe struct{}

func (s *SanAnKe) GetId() string {
	return IdSanAnKe
}

func (s *SanAnKe) GetFan(huWay *ruleset.HuWay) int {
	return 2
}

func (s *SanAnKe) YakumanBaisu() int {
	return 0
}

func (s *SanAnKe) GetName() string {
	return "三暗刻"
}

func (s *SanAnKe) NeedMenQing() bool {
	return false
}

func (s *SanAnKe) GetDescription() string {
	return "胡牌时有三组暗刻，其中包含暗杠也可以。"
}

func (s *SanAnKe) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.GetAnKeCount() >= 3
}

// #endregion

// #region sangangzi

type SanGangZi struct{}

func (s *SanGangZi) GetId() string {
	return IdSanGangZi
}

func (s *SanGangZi) GetFan(huWay *ruleset.HuWay) int {
	return 2
}

func (s *SanGangZi) YakumanBaisu() int {
	return 0
}

func (s *SanGangZi) GetName() string {
	return "三杠子"
}

func (s *SanGangZi) NeedMenQing() bool {
	return false
}

func (s *SanGangZi) GetDescription() string {
	return "胡牌时有三组杠子"
}

func (s *SanGangZi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.Hand.GetGangCount() >= 3
}

// #endregion

// #region qiduizi

type QiDuiZi struct{}

func (q *QiDuiZi) GetId() string {
	return IdQiDuiZi
}

func (q *QiDuiZi) GetFan(huWay *ruleset.HuWay) int {
	return 2
}

func (q *QiDuiZi) YakumanBaisu() int {
	return 0
}

func (q *QiDuiZi) GetName() string {
	return "七对子"
}

func (q *QiDuiZi) NeedMenQing() bool {
	return true
}

func (q *QiDuiZi) GetDescription() string {
	return "由7个对子组成的和牌"
}

func (q *QiDuiZi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsQiDui
}

// #endregion

// #region hunquandaiyao

type HunQuanDaiYao struct{}

func (h *HunQuanDaiYao) GetId() string {
	return IdHunQuanDaiYao
}

func (h *HunQuanDaiYao) GetFan(huWay *ruleset.HuWay) int {
	return GetFuLuMinusOneFan(2, huWay)
}

func (h *HunQuanDaiYao) YakumanBaisu() int {
	return 0
}

func (h *HunQuanDaiYao) GetName() string {
	return "混全带幺九"
}

func (h *HunQuanDaiYao) NeedMenQing() bool {
	return false
}

func (h *HunQuanDaiYao) GetDescription() string {
	return "所有顺子、刻子、杠子、雀头都包含幺九牌。"
}

func (h *HunQuanDaiYao) IsYiZhong(huWay *ruleset.HuWay) bool {
	for _, shunZi := range huWay.GetAllShunZi() {
		if shunZi.Number != 1 && shunZi.Number != 7 {
			return false
		}
	}
	hasZi := false
	for _, keZi := range huWay.GetAllKeZi() {
		if !keZi.IsYaoJiu() {
			return false
		}
		if keZi.IsZi() {
			hasZi = true
		}
	}
	if huWay.QueTou.IsZi() {
		hasZi = true
	}
	return huWay.QueTou.IsYaoJiu() && hasZi
}

// #endregion

// #region hunlaotou

type HunLaoTou struct{}

func (h *HunLaoTou) GetId() string {
	return IdHunLaoTou
}

func (h *HunLaoTou) GetFan(huWay *ruleset.HuWay) int {
	return 2
}

func (h *HunLaoTou) YakumanBaisu() int {
	return 0
}

func (h *HunLaoTou) GetName() string {
	return "混老头"
}

func (h *HunLaoTou) NeedMenQing() bool {
	return false
}

func (h *HunLaoTou) GetDescription() string {
	return "整副牌型只有幺九牌组成"
}

func (h *HunLaoTou) IsYiZhong(huWay *ruleset.HuWay) bool {
	if huWay.GetShunZiCount() > 0 {
		return false
	}
	for _, keZi := range huWay.GetAllKeZi() {
		if !keZi.IsYaoJiu() {
			return false
		}
	}
	return huWay.QueTou.IsYaoJiu()
}

// #endregion

// #region xiaosanyuan

type XiaoSanYuan struct{}

func (x *XiaoSanYuan) GetId() string {
	return IdXiaoSanYuan
}

func (x *XiaoSanYuan) GetFan(huWay *ruleset.HuWay) int {
	return 2
}

func (x *XiaoSanYuan) YakumanBaisu() int {
	return 0
}

func (x *XiaoSanYuan) GetName() string {
	return "小三元"
}

func (x *XiaoSanYuan) NeedMenQing() bool {
	return false
}

func (x *XiaoSanYuan) GetDescription() string {
	return "其中两组三元牌为刻子或杠子，另外一组为对子"
}

func (x *XiaoSanYuan) IsYiZhong(huWay *ruleset.HuWay) bool {
	cnt := 0
	for _, keZi := range huWay.GetAllKeZi() {
		if keZi.IsSanYuan() {
			cnt++
		}
	}
	return cnt == 2 && huWay.QueTou.IsSanYuan()
}

// #endregion

// #region doubleriichi

type DoubleRiichi struct{}

func (d *DoubleRiichi) GetId() string {
	return IdDoubleRiichi
}

func (d *DoubleRiichi) GetFan(huWay *ruleset.HuWay) int {
	return 2
}

func (d *DoubleRiichi) YakumanBaisu() int {
	return 0
}

func (d *DoubleRiichi) GetName() string {
	return "两立直"
}

func (d *DoubleRiichi) NeedMenQing() bool {
	return true
}

func (d *DoubleRiichi) GetDescription() string {
	return "第一巡牌时即宣告“立直”，非第一巡当做一般“立直”处理"
}

func (d *DoubleRiichi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.IsRiichi() && huWay.Hand.DrawNumber == 1
}

// #endregion

// #region hunyise

type HunYiSe struct{}

func (h *HunYiSe) GetId() string {
	return IdHunYiSe
}

func (h *HunYiSe) GetFan(huWay *ruleset.HuWay) int {
	return GetFuLuMinusOneFan(3, huWay)
}

func (h *HunYiSe) YakumanBaisu() int {
	return 0
}

func (h *HunYiSe) GetName() string {
	return "混一色"
}

func (h *HunYiSe) NeedMenQing() bool {
	return false
}

func (h *HunYiSe) GetDescription() string {
	return "由一种花色的牌及字牌组成的和牌"
}

func (h *HunYiSe) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.Hand.IsHunYiSe()
}

// #endregion

// #region chunquandaiyao

type ChunQuanDaiYao struct{}

func (c *ChunQuanDaiYao) GetId() string {
	return IdChunQuanDaiYao
}

func (c *ChunQuanDaiYao) GetFan(huWay *ruleset.HuWay) int {
	return GetFuLuMinusOneFan(3, huWay)
}

func (c *ChunQuanDaiYao) YakumanBaisu() int {
	return 0
}

func (c *ChunQuanDaiYao) GetName() string {
	return "纯全带幺九"
}

func (c *ChunQuanDaiYao) NeedMenQing() bool {
	return false
}

func (c *ChunQuanDaiYao) GetDescription() string {
	return "整个牌型所有顺子、刻子、杠子、对子都包含幺九牌，且无字牌"
}

func (c *ChunQuanDaiYao) IsYiZhong(huWay *ruleset.HuWay) bool {
	for _, shunZi := range huWay.GetAllShunZi() {
		if shunZi.Number != 1 && shunZi.Number != 7 {
			return false
		}
	}
	for _, keZi := range huWay.GetAllKeZi() {
		if keZi.IsZi() || (keZi.Number != 1 && keZi.Number != 9) {
			return false
		}
	}
	return !huWay.QueTou.IsZi() && (huWay.QueTou.Number == 1 || huWay.QueTou.Number == 9)
}

// #endregion

// #region erbeikou

type ErBeiKou struct{}

func (e *ErBeiKou) GetId() string {
	return IdErBeiKou
}

func (e *ErBeiKou) GetFan(huWay *ruleset.HuWay) int {
	return 3
}

func (e *ErBeiKou) YakumanBaisu() int {
	return 0
}

func (e *ErBeiKou) GetName() string {
	return "二盃口"
}

func (e *ErBeiKou) NeedMenQing() bool {
	return true
}

func (e *ErBeiKou) GetDescription() string {
	return "两组相同的顺子"
}

func (e *ErBeiKou) IsYiZhong(huWay *ruleset.HuWay) bool {
	dup := make(map[int]int)
	for _, shunZi := range huWay.Shunzi {
		dup[shunZi.Id]++
	}
	total := 0
	for _, cnt := range dup {
		if cnt >= 2 {
			total += cnt / 2
		}
	}
	return total == 2
}

// #endregion

// #region qingyise

type QingYiSe struct{}

func (q *QingYiSe) GetId() string {
	return IdQingYiSe
}

func (q *QingYiSe) GetFan(huWay *ruleset.HuWay) int {
	return GetFuLuMinusOneFan(6, huWay)
}

func (q *QingYiSe) YakumanBaisu() int {
	return 0
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

func (s *ShiSanYao) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (s *ShiSanYao) YakumanBaisu() int {
	return 1
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

func (s *ShiSanYaoDanJi) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (s *ShiSanYaoDanJi) YakumanBaisu() int {
	return 2
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

// #region dasanyuan

type DaSanYuan struct{}

func (d *DaSanYuan) GetId() string {
	return IdDaSanYuan
}

func (d *DaSanYuan) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (d *DaSanYuan) YakumanBaisu() int {
	return 1
}

func (d *DaSanYuan) GetName() string {
	return "大三元"
}

func (d *DaSanYuan) NeedMenQing() bool {
	return false
}

func (d *DaSanYuan) GetDescription() string {
	return "全数3组三元牌为刻子或杠子"
}

func (d *DaSanYuan) IsYiZhong(huWay *ruleset.HuWay) bool {
	cnt := 0
	for _, keZi := range huWay.GetAllKeZi() {
		if keZi.IsSanYuan() {
			cnt++
		}
	}
	return cnt == 3
}

// #endregion

// #region sianke

type SiAnKe struct{}

func (s *SiAnKe) GetId() string {
	return IdSiAnKe
}

func (s *SiAnKe) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (s *SiAnKe) YakumanBaisu() int {
	return 1
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
	return huWay.GetAnKeCount() == 4 && huWay.IsGotTileQueTou()
}

// #endregion

// #region siankedanji

type SiAnKeDanJi struct{}

func (s *SiAnKeDanJi) GetId() string {
	return IdSiAnKe
}

func (s *SiAnKeDanJi) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (s *SiAnKeDanJi) YakumanBaisu() int {
	return 2
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

func (z *ZiYiSe) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (z *ZiYiSe) YakumanBaisu() int {
	return 1
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

// #region lvyise

type LvYiSe struct{}

func (l *LvYiSe) GetId() string {
	return IdLvYiSe
}

func (l *LvYiSe) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (l *LvYiSe) YakumanBaisu() int {
	return 1
}

func (l *LvYiSe) GetName() string {
	return "绿一色"
}

func (l *LvYiSe) NeedMenQing() bool {
	return false
}

func (l *LvYiSe) GetDescription() string {
	return "由23468条及发组成的和牌"
}

func (l *LvYiSe) isTileGreen(tile *tiles.Tile) bool {
	return tile.Equal(tiles.Fa) ||
		(tile.TileType == tiles.Suo &&
			(tile.Number == 2 || tile.Number == 3 || tile.Number == 4 || tile.Number == 6 || tile.Number == 8))
}

func (l *LvYiSe) IsYiZhong(huWay *ruleset.HuWay) bool {
	for _, tile := range huWay.GetAllShunZi() {
		if !tile.Equal(tiles.Suo2) {
			return false
		}
	}
	for _, tile := range huWay.GetAllKeZi() {
		if !l.isTileGreen(&tile) {
			return false
		}
	}
	return l.isTileGreen(huWay.QueTou)
}

// #endregion

// #region xiaosixi

type XiaoSiXi struct{}

func (x *XiaoSiXi) GetId() string {
	return IdXiaoSiXi
}

func (x *XiaoSiXi) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (x *XiaoSiXi) YakumanBaisu() int {
	return 1
}

func (x *XiaoSiXi) GetName() string {
	return "小四喜"
}

func (x *XiaoSiXi) NeedMenQing() bool {
	return false
}

func (x *XiaoSiXi) GetDescription() string {
	return "3组风牌为刻子或杠子，另1组为对子"
}

func (x *XiaoSiXi) IsYiZhong(huWay *ruleset.HuWay) bool {
	cnt := 0
	for _, keZi := range huWay.GetAllKeZi() {
		if keZi.IsFeng() {
			cnt++
		}
	}
	return cnt == 3 && huWay.QueTou.IsFeng()
}

// #endregion

// #region dasixi

type DaSiXi struct{}

func (d *DaSiXi) GetId() string {
	return IdDaSiXi
}

func (d *DaSiXi) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (d *DaSiXi) YakumanBaisu() int {
	return 2
}

func (d *DaSiXi) GetName() string {
	return "大四喜"
}

func (d *DaSiXi) NeedMenQing() bool {
	return false
}

func (d *DaSiXi) GetDescription() string {
	return "全数4组风牌为刻子或杠子"
}

func (d *DaSiXi) IsYiZhong(huWay *ruleset.HuWay) bool {
	cnt := 0
	for _, keZi := range huWay.GetAllKeZi() {
		if keZi.IsFeng() {
			cnt++
		}
	}
	return cnt == 4
}

// #endregion

// #region qinglaotou

type QingLaoTou struct{}

func (q *QingLaoTou) GetId() string {
	return IdQingLaoTou
}

func (q *QingLaoTou) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (q *QingLaoTou) YakumanBaisu() int {
	return 1
}

func (q *QingLaoTou) GetName() string {
	return "清老头"
}

func (q *QingLaoTou) NeedMenQing() bool {
	return false
}

func (q *QingLaoTou) GetDescription() string {
	return "由序数牌1、9组成的和牌"
}

func (q *QingLaoTou) IsYiZhong(huWay *ruleset.HuWay) bool {
	if huWay.GetShunZiCount() > 0 {
		return false
	}
	for _, tile := range huWay.GetAllKeZi() {
		if tile.IsZi() || (tile.Number != 1 && tile.Number != 9) {
			return false
		}
	}
	return !huWay.QueTou.IsZi() && (huWay.QueTou.Number == 1 || huWay.QueTou.Number == 9)
}

// #endregion

// #region jiulianbaodeng

type JiuLianBaoDeng struct{}

func (j *JiuLianBaoDeng) GetId() string {
	return IdJiuLianBaoDeng
}

func (j *JiuLianBaoDeng) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (j *JiuLianBaoDeng) YakumanBaisu() int {
	return 1
}

func (j *JiuLianBaoDeng) GetName() string {
	return "九莲宝灯"
}

func (j *JiuLianBaoDeng) NeedMenQing() bool {
	return true
}

func (j *JiuLianBaoDeng) GetDescription() string {
	return "在万、筒、条中的某一种以111 2345678 999的形式加上同花色的其中任意一张可成立，且必须门清（暗杠亦不可）"
}

type jiuLianBaoDentType int

const (
	jiuLianBaoDentTypeNone jiuLianBaoDentType = iota
	jiuLianBaoDentTypeNormal
	jiuLianBaoDentTypePure
)

func isTwoToEightAllOneHelper(tileCount *[9]int) bool {
	for i := 1; i < 8; i++ {
		if tileCount[i] != 1 {
			return false
		}
	}
	return true
}

func isPureJiuLianBaoDengHelper(tileCount *[9]int) bool {
	if tileCount[0] == 3 && tileCount[8] == 3 {
		return isTwoToEightAllOneHelper(tileCount)
	}
	return false
}

func getJiuLianBaoDengType(huWay *ruleset.HuWay) jiuLianBaoDentType {
	if len(huWay.Hand.Fulu) > 0 {
		return jiuLianBaoDentTypeNone
	}
	if !huWay.IsQingYiSe() {
		return jiuLianBaoDentTypeNone
	}
	tileCount := [9]int{}
	for _, t := range huWay.Hand.Tiles {
		if t.IsZi() {
			return jiuLianBaoDentTypeNone
		}
		tileCount[t.Number-1]++
	}
	if isPureJiuLianBaoDengHelper(&tileCount) {
		return jiuLianBaoDentTypePure
	}
	tileCount[huWay.GotTile.Tile.Number-1]++
	if tileCount[0] == 4 {
		if tileCount[8] != 3 {
			return jiuLianBaoDentTypeNone
		}
		if isTwoToEightAllOneHelper(&tileCount) {
			return jiuLianBaoDentTypeNormal
		}
		return jiuLianBaoDentTypeNone
	}
	if tileCount[8] == 4 {
		if tileCount[0] != 3 {
			return jiuLianBaoDentTypeNone
		}
		if isTwoToEightAllOneHelper(&tileCount) {
			return jiuLianBaoDentTypeNormal
		}
		return jiuLianBaoDentTypeNone
	}
	if tileCount[0] != 3 || tileCount[8] != 3 {
		return jiuLianBaoDentTypeNone
	}
	tileCount[huWay.GotTile.Tile.Number-1]--
	if isTwoToEightAllOneHelper(&tileCount) {
		return jiuLianBaoDentTypeNormal
	}
	return jiuLianBaoDentTypeNone
}

func (j *JiuLianBaoDeng) IsYiZhong(huWay *ruleset.HuWay) bool {
	return getJiuLianBaoDengType(huWay) == jiuLianBaoDentTypeNormal
}

// #endregion

// #region chunjiulian

type ChunJiuLian struct{}

func (c *ChunJiuLian) GetId() string {
	return IdChunJiuLian
}

func (c *ChunJiuLian) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (c *ChunJiuLian) YakumanBaisu() int {
	return 2
}

func (c *ChunJiuLian) GetName() string {
	return "纯正九莲宝灯"
}

func (c *ChunJiuLian) NeedMenQing() bool {
	return true
}

func (c *ChunJiuLian) GetDescription() string {
	return "九莲宝灯听九面"
}

func (c *ChunJiuLian) IsYiZhong(huWay *ruleset.HuWay) bool {
	return getJiuLianBaoDengType(huWay) == jiuLianBaoDentTypePure
}

// #endregion

// #region sigangzi

type SiGangZi struct{}

func (s *SiGangZi) GetId() string {
	return IdSiGangZi
}

func (s *SiGangZi) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (s *SiGangZi) YakumanBaisu() int {
	return 1
}

func (s *SiGangZi) GetName() string {
	return "四杠子"
}

func (s *SiGangZi) NeedMenQing() bool {
	return false
}

func (s *SiGangZi) GetDescription() string {
	return "由累计四个杠子（明杠和暗杠均可）和一对雀头组成"
}

func (s *SiGangZi) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.Hand.GetGangCount() == 4
}

// #endregion

// #region tianhu

type TianHu struct{}

func (t *TianHu) GetId() string {
	return IdTianHu
}

func (t *TianHu) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (t *TianHu) YakumanBaisu() int {
	return 1
}

func (t *TianHu) GetName() string {
	return "天和"
}

func (t *TianHu) NeedMenQing() bool {
	return true
}

func (t *TianHu) GetDescription() string {
	return "牌局开始时，庄家便自摸。如果是利用最初的14张牌开杠之后岭上开花的话则不成立"
}

func (t *TianHu) IsYiZhong(huWay *ruleset.HuWay) bool {
	return huWay.Hand.Player.IsZhuangJia() && huWay.Hand.DrawNumber == 1 && huWay.IsZiMo() && !huWay.GotTile.IsLingShang
}

// #endregion

// #region dihu

type DiHu struct{}

func (d *DiHu) GetId() string {
	return IdDiHu
}

func (d *DiHu) GetFan(_ *ruleset.HuWay) int {
	return 13
}

func (d *DiHu) YakumanBaisu() int {
	return 1
}

func (d *DiHu) GetName() string {
	return "地和"
}

func (d *DiHu) NeedMenQing() bool {
	return true
}

func (d *DiHu) GetDescription() string {
	return "牌局开始时，闲家第一巡便自摸。如果在第一次自摸之前有任意玩家鸣牌，则不成立"
}

func (d *DiHu) IsYiZhong(huWay *ruleset.HuWay) bool {
	return !huWay.Hand.Player.IsZhuangJia() && !huWay.Hand.Game.HasMingPai && huWay.Hand.DrawNumber == 1 && huWay.IsZiMo() && !huWay.GotTile.IsLingShang
}

// #endregion
