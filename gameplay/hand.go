package gameplay

import (
	"fmt"

	"github.com/0x5ab/gomj/tiles"
)

type FuluType int

const (
	Chi FuluType = iota
	Peng
	Gang
)

type Fulu struct {
	Player     *Player // the source player of this fulu
	Type       FuluType
	StartTile  tiles.Tile
	PlayedTile GameTile
}

func (f *Fulu) IsChi() bool {
	return f.Type == Chi
}

func (f *Fulu) IsPeng() bool {
	return f.Type == Peng
}

func (f *Fulu) IsGang() bool {
	return f.Type == Gang
}

func (f *Fulu) IsAnGang() bool {
	return f.IsGang() && f.PlayedTile.Player.Id == f.Player.Id
}

func (f *Fulu) GetTiles() []tiles.Tile {
	switch f.Type {
	case Chi:
		return []tiles.Tile{f.StartTile, f.StartTile.Next(), f.StartTile.Next().Next()}
	case Peng:
		return []tiles.Tile{f.StartTile, f.StartTile, f.StartTile}
	case Gang:
		return []tiles.Tile{f.StartTile, f.StartTile, f.StartTile, f.StartTile}
	}
	return nil
}

func (f *Fulu) String() string {
	switch f.Type {
	case Chi:
		return fmt.Sprintf("chi %s%s%s", f.StartTile.String(), f.StartTile.Next().String(), f.StartTile.Next().Next().String())
	case Peng:
		return fmt.Sprintf("peng %s", f.StartTile.String())
	case Gang:
		return fmt.Sprintf("kang %s", f.StartTile.String())
	}
	return ""
}

func (f *Fulu) HumanReadableString() string {
	switch f.Type {
	case Chi:
		return fmt.Sprintf("吃%s从%s", f.PlayedTile.Tile.HumanReadableString(), f.StartTile.HumanReadableString())
	case Peng:
		return fmt.Sprintf("碰%s", f.StartTile.HumanReadableString())
	case Gang:
		return fmt.Sprintf("杠%s", f.StartTile.HumanReadableString())
	}
	return ""
}

type Hand struct {
	Player           *Player // the owner of this hand
	Tiles            []tiles.Tile
	Fulu             []Fulu
	IsRiichi         bool
	DrawNumber       int
	DrawsAfterRiichi int
}

func (h *Hand) String() string {
	return fmt.Sprintf("%v %v", h.Tiles, h.Fulu)
}

func (h *Hand) GetPengCount() int {
	count := 0
	for _, f := range h.Fulu {
		if f.IsPeng() {
			count++
		}
	}
	return count
}

func (h *Hand) GetGangCount() int {
	count := 0
	for _, f := range h.Fulu {
		if f.IsGang() {
			count++
		}
	}
	return count
}

func (h *Hand) GetAnGangCount() int {
	count := 0
	for _, f := range h.Fulu {
		if f.IsAnGang() {
			count++
		}
	}
	return count
}

func (h *Hand) GetChiCount() int {
	count := 0
	for _, f := range h.Fulu {
		if f.IsChi() {
			count++
		}
	}
	return count
}

func (h *Hand) GetFuluCount() int {
	return h.GetPengCount() + h.GetGangCount() + h.GetChiCount()
}

func (h *Hand) IsMenQing() bool {
	return h.GetFuluCount() == 0
}

func (h *Hand) IsQingYiSe() bool {
	color := h.Tiles[0].TileType
	for _, t := range h.Tiles {
		if t.TileType != color {
			return false
		}
	}
	for _, f := range h.Fulu {
		if f.StartTile.TileType != color {
			return false
		}
	}
	return true
}

func (h *Hand) IsZiYiSe() bool {
	for _, t := range h.Tiles {
		if !t.IsZi() {
			return false
		}
	}
	for _, f := range h.Fulu {
		if !f.IsPeng() || !f.StartTile.IsZi() {
			return false
		}
	}
	return true
}

func (h *Hand) IsTanYao() bool {
	for _, t := range h.Tiles {
		if t.IsYaoJiu() {
			return false
		}
	}
	return true
}
