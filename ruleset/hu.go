package ruleset

import (
	"strings"

	"github.com/0x5ab/gomj/gameplay"
	"github.com/0x5ab/gomj/tiles"
)

type handForCalc struct {
	Wans  [9]int
	Tongs [9]int
	Suos  [9]int
	Zis   [7]int
}

func newHandForCalc() handForCalc {
	return handForCalc{
		Wans:  [9]int{},
		Tongs: [9]int{},
		Suos:  [9]int{},
		Zis:   [7]int{},
	}
}

func (h *handForCalc) add(tile tiles.Tile) {
	switch tile.TileType {
	case tiles.Wan:
		h.Wans[tile.Number-1]++
	case tiles.Tong:
		h.Tongs[tile.Number-1]++
	case tiles.Suo:
		h.Suos[tile.Number-1]++
	case tiles.Zi:
		h.Zis[tile.Number-1]++
	}
}

func (h *handForCalc) remove(tile tiles.Tile) {
	switch tile.TileType {
	case tiles.Wan:
		h.Wans[tile.Number-1]--
	case tiles.Tong:
		h.Tongs[tile.Number-1]--
	case tiles.Suo:
		h.Suos[tile.Number-1]--
	case tiles.Zi:
		h.Zis[tile.Number-1]--
	}
}

func (h *handForCalc) isNonDivisible(queTouAlreadyTaken bool) bool {
	max := 1
	if queTouAlreadyTaken {
		max = 2
	}
	for _, wan := range h.Wans {
		if wan > max {
			return false
		}
	}
	for _, tong := range h.Tongs {
		if tong > max {
			return false
		}
	}
	for _, suo := range h.Suos {
		if suo > max {
			return false
		}
	}
	for _, zi := range h.Zis {
		if zi > max {
			return false
		}
	}
	return true
}

func (h *handForCalc) isHu() bool {
	for _, wan := range h.Wans {
		if wan != 0 {
			return false
		}
	}
	for _, tong := range h.Tongs {
		if tong != 0 {
			return false
		}
	}
	for _, suo := range h.Suos {
		if suo != 0 {
			return false
		}
	}
	for _, zi := range h.Zis {
		if zi != 0 {
			return false
		}
	}
	return true
}

func (h *handForCalc) isQiDui() bool {
	for _, wan := range h.Wans {
		if wan%2 != 0 {
			return false
		}
	}
	for _, tong := range h.Tongs {
		if tong%2 != 0 {
			return false
		}
	}
	for _, suo := range h.Suos {
		if suo%2 != 0 {
			return false
		}
	}
	for _, zi := range h.Zis {
		if zi%2 != 0 {
			return false
		}
	}
	return true
}

func (h *handForCalc) hasAllYaoJiuTiles() bool {
	return (h.Wans[0] == 1 && h.Wans[8] == 1 &&
		h.Tongs[0] == 1 && h.Tongs[8] == 1 &&
		h.Suos[0] == 1 && h.Suos[8] == 1 &&
		h.Zis[0] == 1 && h.Zis[1] == 1 &&
		h.Zis[2] == 1 && h.Zis[3] == 1 &&
		h.Zis[4] == 1 && h.Zis[5] == 1 &&
		h.Zis[6] == 1)
}

func handToHandForCalc(hand *gameplay.Hand, tile tiles.Tile) handForCalc {
	handForCalc := newHandForCalc()
	for _, t := range hand.Tiles {
		handForCalc.add(t)
	}
	handForCalc.add(tile)
	return handForCalc
}

type HuWay struct {
	isValid          bool
	IsQiDui          bool
	IsShiSanYao      bool
	IsShiSanYaoDanJi bool
	Shunzi           []tiles.Tile
	Kezi             []tiles.Tile
	QueTou           *tiles.Tile
	YiZhongs         []YiZhong
	Point            int
	Fan              int
	Hand             *gameplay.Hand
	GotTile          *gameplay.GameTile
}

func NewHuWay(hand *gameplay.Hand, tile *gameplay.GameTile) HuWay {
	return HuWay{
		Hand:    hand,
		GotTile: tile,
	}
}

func (h *HuWay) Clone() HuWay {
	clone := HuWay{
		isValid:          h.isValid,
		IsQiDui:          h.IsQiDui,
		IsShiSanYao:      h.IsShiSanYao,
		IsShiSanYaoDanJi: h.IsShiSanYaoDanJi,
		Shunzi:           make([]tiles.Tile, len(h.Shunzi)),
		Kezi:             make([]tiles.Tile, len(h.Kezi)),
		YiZhongs:         make([]YiZhong, len(h.YiZhongs)),
		Point:            h.Point,
		Fan:              h.Fan,
		Hand:             h.Hand,
		GotTile:          h.GotTile,
	}
	copy(clone.Shunzi, h.Shunzi)
	copy(clone.Kezi, h.Kezi)
	copy(clone.YiZhongs, h.YiZhongs)
	if h.QueTou != nil {
		tile := *h.QueTou
		clone.QueTou = &tile
	}
	return clone
}

func (h *HuWay) GetKeZiCount() int {
	return h.GetKeZiCountWithoutPeng() + h.Hand.GetPengCount() + h.Hand.GetGangCount()
}

func (h *HuWay) GetKeZiCountWithoutPeng() int {
	return len(h.Kezi)
}

func (h *HuWay) GetAllKeZi() []tiles.Tile {
	keZi := make([]tiles.Tile, 0, h.GetKeZiCount())
	keZi = append(keZi, h.Hand.GetKeZiTiles()...)
	keZi = append(keZi, h.Kezi...)
	return keZi
}

func (h *HuWay) GetShunZiCount() int {
	return h.GetShunZiCountWithoutChi() + h.Hand.GetChiCount()
}

func (h *HuWay) GetShunZiCountWithoutChi() int {
	return len(h.Shunzi)
}

func (h *HuWay) GetAllShunZi() []tiles.Tile {
	shunZi := make([]tiles.Tile, 0, h.GetShunZiCount())
	shunZi = append(shunZi, h.Hand.GetChiTiles()...)
	shunZi = append(shunZi, h.Shunzi...)
	return shunZi
}

func (h *HuWay) IsValid() bool {
	return h.isValid
}

func (h *HuWay) IsMenQing() bool {
	return h.Hand.IsMenQing()
}

func (h *HuWay) IsZiMo() bool {
	return h.GotTile.State() == gameplay.GameTileDrawn
}

func (h *HuWay) IsRiichi() bool {
	return h.Hand.IsRiichi
}

func (h *HuWay) IsTanYao() bool {
	return h.Hand.IsTanYao() && !h.GotTile.Tile.IsYaoJiu()
}

func (h *HuWay) IsQingYiSe() bool {
	return h.Hand.IsQingYiSe()
}

// IsGotTileQueTou returns true if the got tile is the same as the que tou (is 单钓).
func (h *HuWay) IsGotTileQueTou() bool {
	return h.IsQiDui || (h.QueTou != nil && h.QueTou.Equal(h.GotTile.Tile))
}

func (h *HuWay) HumanReadableString() string {
	if h.IsQiDui {
		return "七对"
	}
	if h.IsShiSanYao {
		if h.IsShiSanYaoDanJi {
			return "十三幺单骑"
		}
		return "十三幺"
	}
	s := strings.Builder{}
	if len(h.Shunzi) > 0 {
		s.WriteString("顺子：")
		for _, shunzi := range h.Shunzi {
			s.WriteString(shunzi.HumanReadableString())
			s.WriteString("，")
		}
	}
	s.WriteString(" | ")
	if len(h.Kezi) > 0 {
		s.WriteString("刻子：")
		for _, kezi := range h.Kezi {
			s.WriteString(kezi.HumanReadableString())
			s.WriteString("，")
		}
	}
	s.WriteString(" | ")
	if h.QueTou != nil {
		s.WriteString("雀头：")
		s.WriteString(h.QueTou.HumanReadableString())
	}
	return s.String()
}

func (h *HuWay) CalculateStats(r *Ruleset) {
	h.YiZhongs = r.Check(h)
	h.Point = r.GetPoint(h)
	h.Fan = 0
	for _, yizhong := range h.YiZhongs {
		h.Fan += yizhong.GetFan(h.Hand)
	}
}

type TingPaiType int

const (
	TingPaiTypeNone TingPaiType = iota
	TingPaiTypeNormal
	TingPaiTypeQiDui
	TingPaiTypeShiSanYao
	// 单钓，边张，或者嵌张
	TingPaiTypeSpecial
)

func (h *HuWay) GetTingPaiType() TingPaiType {
	if !h.isValid {
		return TingPaiTypeNone
	}
	if h.IsQiDui {
		return TingPaiTypeQiDui
	}
	if h.IsShiSanYao {
		return TingPaiTypeShiSanYao
	}
	// check if the got tile is a 单钓
	if h.IsGotTileQueTou() {
		return TingPaiTypeSpecial
	}
	// check if the got tile is a 边张
	if !h.GotTile.Tile.IsZi() &&
		(h.GotTile.Tile.Number == 3 || h.GotTile.Tile.Number == 7) {
		for _, shunzi := range h.Shunzi {
			if shunzi.Number == 1 || shunzi.Number == 7 {
				return TingPaiTypeSpecial
			}
		}
	}
	// check if the got tile is a 嵌张
	if !h.GotTile.Tile.IsZi() &&
		(h.GotTile.Tile.Number == 2 || h.GotTile.Tile.Number == 8) {
		for _, shunzi := range h.Shunzi {
			if shunzi.Number == 3 || shunzi.Number == 7 {
				return TingPaiTypeSpecial
			}
		}
	}
	return TingPaiTypeNormal
}

func isShiSanYao(handForCalc handForCalc, tile tiles.Tile) *tiles.Tile {
	queTou := tiles.Invalid
	if handForCalc.Wans[0] == 2 {
		queTou = tiles.Wan1
	} else if handForCalc.Wans[8] == 2 {
		queTou = tiles.Wan9
	} else if handForCalc.Tongs[0] == 2 {
		queTou = tiles.Tong1
	} else if handForCalc.Tongs[8] == 2 {
		queTou = tiles.Tong9
	} else if handForCalc.Suos[0] == 2 {
		queTou = tiles.Suo1
	} else if handForCalc.Suos[8] == 2 {
		queTou = tiles.Suo9
	} else if handForCalc.Zis[0] == 2 {
		queTou = tiles.Dong
	} else if handForCalc.Zis[1] == 2 {
		queTou = tiles.Nan
	} else if handForCalc.Zis[2] == 2 {
		queTou = tiles.Xi
	} else if handForCalc.Zis[3] == 2 {
		queTou = tiles.Bei
	} else if handForCalc.Zis[4] == 2 {
		queTou = tiles.Zhong
	} else if handForCalc.Zis[5] == 2 {
		queTou = tiles.Fa
	} else if handForCalc.Zis[6] == 2 {
		queTou = tiles.Bai
	}
	if queTou == tiles.Invalid {
		return nil
	}
	handForCalc.remove(queTou)
	if handForCalc.hasAllYaoJiuTiles() {
		return &queTou
	}
	return nil
}

func isShiSanYaoDanJi(handForCalc handForCalc, tile tiles.Tile) bool {
	handForCalc.remove(tile)
	return handForCalc.hasAllYaoJiuTiles()
}

func CanHu(ruleset *Ruleset, hand *gameplay.Hand, tile *gameplay.GameTile) *HuWay {
	handForCalc := handToHandForCalc(hand, tile.Tile)
	huWay := NewHuWay(hand, tile)
	if handForCalc.isQiDui() {
		huWay.isValid = true
		huWay.IsQiDui = true
		huWay.CalculateStats(ruleset)
		return &huWay
	}
	if isShiSanYaoDanJi(handForCalc, tile.Tile) {
		// huWays.Ways = getShiSanYaoDanJiWays()
		huWay.isValid = true
		huWay.IsShiSanYao = true
		huWay.IsShiSanYaoDanJi = true
		huWay.CalculateStats(ruleset)
		return &huWay
	}
	if queTou := isShiSanYao(handForCalc, tile.Tile); queTou != nil {
		huWay.isValid = true
		huWay.IsShiSanYao = true
		huWay.QueTou = queTou
		huWay.CalculateStats(ruleset)
		return &huWay
	}
	result := NewHuWay(hand, tile)
	ruleset.canHu(&result, &huWay, &handForCalc)
	return &result
}

func (r *Ruleset) canHu(result *HuWay, currentWay *HuWay, handForCalc *handForCalc) {
	if handForCalc.isNonDivisible(currentWay.QueTou != nil) {
		if handForCalc.isHu() {
			// we need only the hu way with the highest fan and/or points/fu
			currentWay.isValid = true
			currentWay.CalculateStats(r)
			if currentWay.Fan > result.Fan || (currentWay.Fan == result.Fan && currentWay.Point > result.Point) {
				*result = currentWay.Clone()
			}
		}
		return
	}
	for i := 0; i < 9; i++ {
		// Wan
		if handForCalc.Wans[i] >= 3 {
			handForCalc.Wans[i] -= 3
			currentWay.Kezi = append(currentWay.Kezi, tiles.GetTileP(tiles.Wan, i+1))
			r.canHu(result, currentWay, handForCalc)
			currentWay.Kezi = currentWay.Kezi[:len(currentWay.Kezi)-1]
			handForCalc.Wans[i] += 3
		}
		if currentWay.QueTou == nil && handForCalc.Wans[i] >= 2 {
			handForCalc.Wans[i] -= 2
			currentWay.QueTou = tiles.GetTileP(tiles.Wan, i+1).Ptr()
			r.canHu(result, currentWay, handForCalc)
			currentWay.QueTou = nil
			handForCalc.Wans[i] += 2
		}
		if i < 7 && handForCalc.Wans[i] >= 1 && handForCalc.Wans[i+1] >= 1 && handForCalc.Wans[i+2] >= 1 {
			handForCalc.Wans[i] -= 1
			handForCalc.Wans[i+1] -= 1
			handForCalc.Wans[i+2] -= 1
			currentWay.Shunzi = append(currentWay.Shunzi, tiles.GetTileP(tiles.Wan, i+1))
			r.canHu(result, currentWay, handForCalc)
			currentWay.Shunzi = currentWay.Shunzi[:len(currentWay.Shunzi)-1]
			handForCalc.Wans[i] += 1
			handForCalc.Wans[i+1] += 1
			handForCalc.Wans[i+2] += 1
		}
		// Tong
		if handForCalc.Tongs[i] >= 3 {
			handForCalc.Tongs[i] -= 3
			currentWay.Kezi = append(currentWay.Kezi, tiles.GetTileP(tiles.Tong, i+1))
			r.canHu(result, currentWay, handForCalc)
			currentWay.Kezi = currentWay.Kezi[:len(currentWay.Kezi)-1]
			handForCalc.Tongs[i] += 3
		}
		if currentWay.QueTou == nil && handForCalc.Tongs[i] >= 2 {
			handForCalc.Tongs[i] -= 2
			currentWay.QueTou = tiles.GetTileP(tiles.Tong, i+1).Ptr()
			r.canHu(result, currentWay, handForCalc)
			currentWay.QueTou = nil
			handForCalc.Tongs[i] += 2
		}
		if i < 7 && handForCalc.Tongs[i] >= 1 && handForCalc.Tongs[i+1] >= 1 && handForCalc.Tongs[i+2] >= 1 {
			handForCalc.Tongs[i] -= 1
			handForCalc.Tongs[i+1] -= 1
			handForCalc.Tongs[i+2] -= 1
			currentWay.Shunzi = append(currentWay.Shunzi, tiles.GetTileP(tiles.Tong, i+1))
			r.canHu(result, currentWay, handForCalc)
			currentWay.Shunzi = currentWay.Shunzi[:len(currentWay.Shunzi)-1]
			handForCalc.Tongs[i] += 1
			handForCalc.Tongs[i+1] += 1
			handForCalc.Tongs[i+2] += 1
		}
		// Suo
		if handForCalc.Suos[i] >= 3 {
			handForCalc.Suos[i] -= 3
			currentWay.Kezi = append(currentWay.Kezi, tiles.GetTileP(tiles.Suo, i+1))
			r.canHu(result, currentWay, handForCalc)
			currentWay.Kezi = currentWay.Kezi[:len(currentWay.Kezi)-1]
			handForCalc.Suos[i] += 3
		}
		if currentWay.QueTou == nil && handForCalc.Suos[i] >= 2 {
			handForCalc.Suos[i] -= 2
			currentWay.QueTou = tiles.GetTileP(tiles.Suo, i+1).Ptr()
			r.canHu(result, currentWay, handForCalc)
			currentWay.QueTou = nil
			handForCalc.Suos[i] += 2
		}
		if i < 7 && handForCalc.Suos[i] >= 1 && handForCalc.Suos[i+1] >= 1 && handForCalc.Suos[i+2] >= 1 {
			handForCalc.Suos[i] -= 1
			handForCalc.Suos[i+1] -= 1
			handForCalc.Suos[i+2] -= 1
			currentWay.Shunzi = append(currentWay.Shunzi, tiles.GetTileP(tiles.Suo, i+1))
			r.canHu(result, currentWay, handForCalc)
			currentWay.Shunzi = currentWay.Shunzi[:len(currentWay.Shunzi)-1]
			handForCalc.Suos[i] += 1
			handForCalc.Suos[i+1] += 1
			handForCalc.Suos[i+2] += 1
		}
	}
	for i := 0; i < 7; i++ {
		// Zi
		if handForCalc.Zis[i] >= 3 {
			handForCalc.Zis[i] -= 3
			currentWay.Kezi = append(currentWay.Kezi, tiles.GetTileP(tiles.Zi, i+1))
			r.canHu(result, currentWay, handForCalc)
			currentWay.Kezi = currentWay.Kezi[:len(currentWay.Kezi)-1]
			handForCalc.Zis[i] += 3
		}
		if currentWay.QueTou == nil && handForCalc.Zis[i] >= 2 {
			handForCalc.Zis[i] -= 2
			currentWay.QueTou = tiles.GetTileP(tiles.Zi, i+1).Ptr()
			r.canHu(result, currentWay, handForCalc)
			currentWay.QueTou = nil
			handForCalc.Zis[i] += 2
		}
	}
}
