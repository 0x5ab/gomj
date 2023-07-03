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

func (f Fulu) String() string {
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

// TODO: optimize: do not recalculate stats for fulu every time

type Hand struct {
	Game             *Game
	Player           *Player // the owner of this hand
	Tiles            []tiles.Tile
	Fulu             []Fulu
	IsRiichi         bool
	DrawNumber       int
	DrawsAfterRiichi int
}

func NewHand(g *Game, p *Player) *Hand {
	return &Hand{
		Game:   g,
		Player: p,
		Tiles:  []tiles.Tile{},
		Fulu:   []Fulu{},
	}
}

func (h *Hand) String() string {
	return fmt.Sprintf("tiles: %v, fulu: %v", h.Tiles, h.Fulu)
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

func (h *Hand) GetMingGangCount() int {
	count := 0
	for _, f := range h.Fulu {
		if f.IsGang() && !f.IsAnGang() {
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
	return h.GetPengCount() + h.GetMingGangCount() + h.GetChiCount()
}

func (h *Hand) GetPengTiles() []tiles.Tile {
	var tiles []tiles.Tile
	for _, f := range h.Fulu {
		if f.IsPeng() {
			tiles = append(tiles, f.StartTile)
		}
	}
	return tiles
}

func (h *Hand) GetGangTiles() []tiles.Tile {
	var tiles []tiles.Tile
	for _, f := range h.Fulu {
		if f.IsGang() {
			tiles = append(tiles, f.StartTile)
		}
	}
	return tiles
}

func (h *Hand) GetMingGangTiles() []tiles.Tile {
	var tiles []tiles.Tile
	for _, f := range h.Fulu {
		if f.IsGang() && !f.IsAnGang() {
			tiles = append(tiles, f.StartTile)
		}
	}
	return tiles
}

func (h *Hand) GetAnGangTiles() []tiles.Tile {
	var tiles []tiles.Tile
	for _, f := range h.Fulu {
		if f.IsAnGang() {
			tiles = append(tiles, f.StartTile)
		}
	}
	return tiles
}

// GetKeZiTiles returns all tiles representing a peng or gang
func (h *Hand) GetKeZiTiles() []tiles.Tile {
	var tiles []tiles.Tile
	for _, f := range h.Fulu {
		if f.IsPeng() || f.IsGang() {
			tiles = append(tiles, f.StartTile)
		}
	}
	return tiles
}

func (h *Hand) GetChiTiles() []tiles.Tile {
	var tiles []tiles.Tile
	for _, f := range h.Fulu {
		if f.IsChi() {
			tiles = append(tiles, f.StartTile)
		}
	}
	return tiles
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

func (h *Hand) IsHunYiSe() bool {
	var tileType tiles.TileType = 0
	hasZi := false
	for _, t := range h.Tiles {
		if tileType == 0 {
			tileType = t.TileType
		} else if tileType != t.TileType && !t.IsZi() {
			return false
		}
		if t.IsZi() {
			hasZi = true
		}
	}
	for _, f := range h.Fulu {
		if tileType == 0 {
			tileType = f.StartTile.TileType
		} else if tileType != f.StartTile.TileType && !f.StartTile.IsZi() {
			return false
		}
		if f.StartTile.IsZi() {
			hasZi = true
		}
	}
	return tileType != 0 && hasZi
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
