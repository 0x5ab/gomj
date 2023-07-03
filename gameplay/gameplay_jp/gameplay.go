package gameplay_jp

import (
	"github.com/0x5ab/gomj/gameplay"
	"github.com/0x5ab/gomj/tiles"
)

type JpMahjongGameTile gameplay.GameTile

type JpMahjongGameTileAdditionalProps struct {
	IsDora    bool
	IsUraDora bool
	IsAkaDora bool
}

func NewJpMahjongGameTile(t tiles.Tile, player *gameplay.Player) *JpMahjongGameTile {
	return &JpMahjongGameTile{
		Tile:             t,
		SpecialTileProps: &JpMahjongGameTileAdditionalProps{},
	}
}
