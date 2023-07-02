package tiles

import (
	"fmt"

	"github.com/0x5ab/gomj/errors"
	"github.com/0x5ab/gomj/utils"
	"github.com/0x5ab/gomj/wind"
)

type TileType rune

const (
	Wan  TileType = 'm'
	Tong TileType = 'p'
	Suo  TileType = 's'
	Zi   TileType = 'z'
)

type Tile struct {
	Id       int
	TileType TileType
	Number   int
}

var (
	Invalid = Tile{}
	// 万
	Wan1 = Tile{1, Wan, 1}
	Wan2 = Tile{2, Wan, 2}
	Wan3 = Tile{3, Wan, 3}
	Wan4 = Tile{4, Wan, 4}
	Wan5 = Tile{5, Wan, 5}
	Wan6 = Tile{6, Wan, 6}
	Wan7 = Tile{7, Wan, 7}
	Wan8 = Tile{8, Wan, 8}
	Wan9 = Tile{9, Wan, 9}
	// 筒
	Tong1 = Tile{11, Tong, 1}
	Tong2 = Tile{12, Tong, 2}
	Tong3 = Tile{13, Tong, 3}
	Tong4 = Tile{14, Tong, 4}
	Tong5 = Tile{15, Tong, 5}
	Tong6 = Tile{16, Tong, 6}
	Tong7 = Tile{17, Tong, 7}
	Tong8 = Tile{18, Tong, 8}
	Tong9 = Tile{19, Tong, 9}
	// 索
	Suo1 = Tile{21, Suo, 1}
	Suo2 = Tile{22, Suo, 2}
	Suo3 = Tile{23, Suo, 3}
	Suo4 = Tile{24, Suo, 4}
	Suo5 = Tile{25, Suo, 5}
	Suo6 = Tile{26, Suo, 6}
	Suo7 = Tile{27, Suo, 7}
	Suo8 = Tile{28, Suo, 8}
	Suo9 = Tile{29, Suo, 9}
	// 字
	Dong  = Tile{31, Zi, int(wind.East)}
	Nan   = Tile{32, Zi, int(wind.South)}
	Xi    = Tile{33, Zi, int(wind.West)}
	Bei   = Tile{34, Zi, int(wind.North)}
	Zhong = Tile{35, Zi, 5}
	Fa    = Tile{36, Zi, 6}
	Bai   = Tile{37, Zi, 7}
)

var (
	AllTiles = []Tile{
		Wan1, Wan2, Wan3, Wan4, Wan5, Wan6, Wan7, Wan8, Wan9,
		Tong1, Tong2, Tong3, Tong4, Tong5, Tong6, Tong7, Tong8, Tong9,
		Suo1, Suo2, Suo3, Suo4, Suo5, Suo6, Suo7, Suo8, Suo9,
		Dong, Nan, Xi, Bei, Zhong, Fa, Bai,
	}
	Wans  = []Tile{Wan1, Wan2, Wan3, Wan4, Wan5, Wan6, Wan7, Wan8, Wan9}
	Tongs = []Tile{Tong1, Tong2, Tong3, Tong4, Tong5, Tong6, Tong7, Tong8, Tong9}
	Suos  = []Tile{Suo1, Suo2, Suo3, Suo4, Suo5, Suo6, Suo7, Suo8, Suo9}
	Zis   = []Tile{Dong, Nan, Xi, Bei, Zhong, Fa, Bai}
)

func (t Tile) IsValid() bool {
	return t != Invalid
}

func (t Tile) Next() Tile {
	if t.Number == 9 || t.TileType == Zi {
		return t
	}
	return Tile{t.Id + 1, t.TileType, t.Number + 1}
}

func (t Tile) IsYaoJiu() bool {
	return t.Number == 1 || t.Number == 9 || t.TileType == Zi
}

func (t Tile) IsZi() bool {
	return t.TileType == Zi
}

func (t Tile) IsFeng() bool {
	return t.TileType == Zi && t.Number <= 4
}

func (t Tile) IsSanYuan() bool {
	return t.TileType == Zi && t.Number >= 5
}

func (t Tile) Equal(t2 Tile) bool {
	return t.Id == t2.Id
}

func (t Tile) String() string {
	return fmt.Sprintf("%d%c", t.Number, t.TileType)
}

func (t Tile) HumanReadableString() string {
	switch t.TileType {
	case Wan:
		return fmt.Sprintf("%s万", utils.DigitToChinese(t.Number))
	case Tong:
		return fmt.Sprintf("%s筒", utils.DigitToChinese(t.Number))
	case Suo:
		return fmt.Sprintf("%s索", utils.DigitToChinese(t.Number))
	case Zi:
		switch t.Number {
		case Dong.Number:
			return "东"
		case Nan.Number:
			return "南"
		case Xi.Number:
			return "西"
		case Bei.Number:
			return "北"
		case Zhong.Number:
			return "中"
		case Fa.Number:
			return "发"
		case Bai.Number:
			return "白"
		}
	}
	return ""
}

func (t Tile) IsWindType(w wind.Wind) bool {
	return t.TileType == Zi && t.Number == int(w)
}

func (t Tile) Ptr() *Tile {
	return &t
}

func GetTile(t TileType, number int) (Tile, error) {
	if number < 1 || number > 9 {
		return Invalid, errors.ErrInvalidTile
	}
	switch t {
	case Wan:
		return Wans[number-1], nil
	case Tong:
		return Tongs[number-1], nil
	case Suo:
		return Suos[number-1], nil
	case Zi:
		switch number {
		case Dong.Number:
			return Dong, nil
		case Nan.Number:
			return Nan, nil
		case Xi.Number:
			return Xi, nil
		case Bei.Number:
			return Bei, nil
		case Zhong.Number:
			return Zhong, nil
		case Fa.Number:
			return Fa, nil
		case Bai.Number:
			return Bai, nil
		}
	}
	return Invalid, errors.ErrInvalidTile
}

func GetTileP(t TileType, number int) Tile {
	tile, err := GetTile(t, number)
	if err != nil {
		panic(err)
	}
	return tile
}
