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

func handToHandForCalc(hand gameplay.Hand, tile tiles.Tile) handForCalc {
	handForCalc := newHandForCalc()
	for _, t := range hand.Tiles {
		handForCalc.add(t)
	}
	handForCalc.add(tile)
	return handForCalc
}

type HuWay struct {
	IsQiDui          bool
	IsShiSanYao      bool
	IsShiSanYaoDanJi bool
	Shunzi           []tiles.Tile
	Kezi             []tiles.Tile
	QueTou           *tiles.Tile
	YiZhongs         []YiZhong
}

type HuWays struct {
	Ways    []HuWay
	Hand    gameplay.Hand
	GotTile *gameplay.GameTile // const
}

func (h *HuWays) calcYiZhongs(yizhongRuleset YiZhongRuleset) {
	for i := range h.Ways {
		h.Ways[i].YiZhongs = yizhongRuleset.Check(&h.Ways[i], &h.Hand, h.GotTile)
	}
}

func (h *HuWay) Clone() HuWay {
	clone := HuWay{
		IsQiDui:          h.IsQiDui,
		IsShiSanYao:      h.IsShiSanYao,
		IsShiSanYaoDanJi: h.IsShiSanYaoDanJi,
		Shunzi:           make([]tiles.Tile, len(h.Shunzi)),
		Kezi:             make([]tiles.Tile, len(h.Kezi)),
	}
	copy(clone.Shunzi, h.Shunzi)
	copy(clone.Kezi, h.Kezi)
	if h.QueTou != nil {
		tile := *h.QueTou
		clone.QueTou = &tile
	}
	return clone
}

func (h *HuWay) GetKeZiCount() int {
	return len(h.Kezi)
}

func (h *HuWay) GetShunZiCount() int {
	return len(h.Shunzi)
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
	if len(h.Kezi) > 0 {
		s.WriteString("刻子：")
		for _, kezi := range h.Kezi {
			s.WriteString(kezi.HumanReadableString())
			s.WriteString("，")
		}
	}
	if h.QueTou != nil {
		s.WriteString("雀头：")
		s.WriteString(h.QueTou.HumanReadableString())
	}
	return s.String()
}

func hasAllYaoJiuTiles(handForCalc handForCalc) bool {
	return (handForCalc.Wans[0] == 1 && handForCalc.Wans[8] == 1 &&
		handForCalc.Tongs[0] == 1 && handForCalc.Tongs[8] == 1 &&
		handForCalc.Suos[0] == 1 && handForCalc.Suos[8] == 1 &&
		handForCalc.Zis[0] == 1 && handForCalc.Zis[1] == 1 &&
		handForCalc.Zis[2] == 1 && handForCalc.Zis[3] == 1 &&
		handForCalc.Zis[4] == 1 && handForCalc.Zis[5] == 1 &&
		handForCalc.Zis[6] == 1)
}

func isShiSanYao(handForCalc handForCalc, tile tiles.Tile) *HuWay {
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
	if hasAllYaoJiuTiles(handForCalc) {
		return &HuWay{
			IsShiSanYao:      true,
			IsShiSanYaoDanJi: false,
			QueTou:           &queTou,
		}
	}
	return nil
}

func isShiSanYaoDanJi(handForCalc handForCalc, tile tiles.Tile) bool {
	handForCalc.remove(tile)
	return hasAllYaoJiuTiles(handForCalc)
}

func getShiSanYaoDanJiWays() []HuWay {
	ways := make([]HuWay, 13)
	for _, way := range ways {
		way.IsShiSanYao = true
		way.IsShiSanYaoDanJi = true
	}
	ways[0].QueTou = &tiles.Wan1
	ways[1].QueTou = &tiles.Wan9
	ways[2].QueTou = &tiles.Tong1
	ways[3].QueTou = &tiles.Tong9
	ways[4].QueTou = &tiles.Suo1
	ways[5].QueTou = &tiles.Suo9
	for i := 0; i < 7; i++ {
		ways[i+6].QueTou = &tiles.Zis[i]
	}
	return ways
}

func CanHu(ruleset YiZhongRuleset, hand gameplay.Hand, tile *gameplay.GameTile) *HuWays {
	handForCalc := handToHandForCalc(hand, tile.Tile)
	huWays := HuWays{Hand: hand, GotTile: tile}
	if handForCalc.isQiDui() {
		huWays.Ways = append(huWays.Ways, HuWay{IsQiDui: true})
		return &huWays
	}
	if isShiSanYaoDanJi(handForCalc, tile.Tile) {
		huWays.Ways = getShiSanYaoDanJiWays()
		return &huWays
	}
	if way := isShiSanYao(handForCalc, tile.Tile); way != nil {
		huWays.Ways = append(huWays.Ways, *way)
		return &huWays
	}
	currentWay := &HuWay{}
	canHu(&huWays, currentWay, &handForCalc)
	huWays.calcYiZhongs(ruleset)
	return &huWays
}

func canHu(huWays *HuWays, currentWay *HuWay, handForCalc *handForCalc) bool {
	if handForCalc.isNonDivisible(currentWay.QueTou != nil) {
		if handForCalc.isHu() {
			huWays.Ways = append(huWays.Ways, currentWay.Clone())
			return true
		}
		return false
	}
	for i := 0; i < 9; i++ {
		// Wan
		if currentWay.QueTou == nil && handForCalc.Wans[i] >= 2 {
			handForCalc.Wans[i] -= 2
			currentWay.QueTou = &tiles.Tile{TileType: tiles.Wan, Number: i + 1}
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.QueTou = nil
			handForCalc.Wans[i] += 2
		}
		if handForCalc.Wans[i] >= 3 {
			handForCalc.Wans[i] -= 3
			currentWay.Kezi = append(currentWay.Kezi, tiles.Tile{TileType: tiles.Wan, Number: i + 1})
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.Kezi = currentWay.Kezi[:len(currentWay.Kezi)-1]
			handForCalc.Wans[i] += 3
		}
		if i < 7 && handForCalc.Wans[i] >= 1 && handForCalc.Wans[i+1] >= 1 && handForCalc.Wans[i+2] >= 1 {
			handForCalc.Wans[i] -= 1
			handForCalc.Wans[i+1] -= 1
			handForCalc.Wans[i+2] -= 1
			currentWay.Shunzi = append(currentWay.Shunzi, tiles.Tile{TileType: tiles.Wan, Number: i + 1})
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.Shunzi = currentWay.Shunzi[:len(currentWay.Shunzi)-1]
			handForCalc.Wans[i] += 1
			handForCalc.Wans[i+1] += 1
			handForCalc.Wans[i+2] += 1
		}
		// Tong
		if currentWay.QueTou == nil && handForCalc.Tongs[i] >= 2 {
			handForCalc.Tongs[i] -= 2
			currentWay.QueTou = &tiles.Tile{TileType: tiles.Tong, Number: i + 1}
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.QueTou = nil
			handForCalc.Tongs[i] += 2
		}
		if handForCalc.Tongs[i] >= 3 {
			handForCalc.Tongs[i] -= 3
			currentWay.Kezi = append(currentWay.Kezi, tiles.Tile{TileType: tiles.Tong, Number: i + 1})
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.Kezi = currentWay.Kezi[:len(currentWay.Kezi)-1]
			handForCalc.Tongs[i] += 3
		}
		if i < 7 && handForCalc.Tongs[i] >= 1 && handForCalc.Tongs[i+1] >= 1 && handForCalc.Tongs[i+2] >= 1 {
			handForCalc.Tongs[i] -= 1
			handForCalc.Tongs[i+1] -= 1
			handForCalc.Tongs[i+2] -= 1
			currentWay.Shunzi = append(currentWay.Shunzi, tiles.Tile{TileType: tiles.Tong, Number: i + 1})
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.Shunzi = currentWay.Shunzi[:len(currentWay.Shunzi)-1]
			handForCalc.Tongs[i] += 1
			handForCalc.Tongs[i+1] += 1
			handForCalc.Tongs[i+2] += 1
		}
		// Suo
		if currentWay.QueTou == nil && handForCalc.Suos[i] >= 2 {
			handForCalc.Suos[i] -= 2
			currentWay.QueTou = &tiles.Tile{TileType: tiles.Suo, Number: i + 1}
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.QueTou = nil
			handForCalc.Suos[i] += 2
		}
		if handForCalc.Suos[i] >= 3 {
			handForCalc.Suos[i] -= 3
			currentWay.Kezi = append(currentWay.Kezi, tiles.Tile{TileType: tiles.Suo, Number: i + 1})
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.Kezi = currentWay.Kezi[:len(currentWay.Kezi)-1]
			handForCalc.Suos[i] += 3
		}
		if i < 7 && handForCalc.Suos[i] >= 1 && handForCalc.Suos[i+1] >= 1 && handForCalc.Suos[i+2] >= 1 {
			handForCalc.Suos[i] -= 1
			handForCalc.Suos[i+1] -= 1
			handForCalc.Suos[i+2] -= 1
			currentWay.Shunzi = append(currentWay.Shunzi, tiles.Tile{TileType: tiles.Suo, Number: i + 1})
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.Shunzi = currentWay.Shunzi[:len(currentWay.Shunzi)-1]
			handForCalc.Suos[i] += 1
			handForCalc.Suos[i+1] += 1
			handForCalc.Suos[i+2] += 1
		}
	}
	for i := 0; i < 7; i++ {
		// Zi
		if currentWay.QueTou == nil && handForCalc.Zis[i] >= 2 {
			handForCalc.Zis[i] -= 2
			currentWay.QueTou = &tiles.Tile{TileType: tiles.Zi, Number: i + 1}
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.QueTou = nil
			handForCalc.Zis[i] += 2
		}
		if handForCalc.Zis[i] >= 3 {
			handForCalc.Zis[i] -= 3
			currentWay.Kezi = append(currentWay.Kezi, tiles.Tile{TileType: tiles.Zi, Number: i + 1})
			if canHu(huWays, currentWay, handForCalc) {
				return true
			}
			currentWay.Kezi = currentWay.Kezi[:len(currentWay.Kezi)-1]
			handForCalc.Zis[i] += 3
		}
	}
	return false
}
