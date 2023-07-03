package gameplay

import "github.com/0x5ab/gomj/tiles"

type GameTileState int

const (
	GameTileInWall GameTileState = iota
	// just drawed from wall
	GameTileDrawn
	GameTileInHand
	GameTilePlayed
	GameTileNotUsed
)

type GameTile struct {
	IsLingShang      bool // is this tile ling shang
	IsLast           bool // is this last drawable tile
	Tile             tiles.Tile
	Player           *Player       // who played or is owning this tile, nil if it's a tile in the wall
	state            GameTileState // state of this tile
	SpecialTileProps any           // any special rules
}

func (t *GameTile) State() GameTileState {
	return t.state
}

func (t *GameTile) Draw() {
	t.state = GameTileDrawn
}

func (t *GameTile) InHand() {
	t.state = GameTileInHand
}

func (t *GameTile) Play() {
	t.state = GameTilePlayed
}

// IsJustDrawn returns true if this tile is just drawn
func (t *GameTile) IsJustDrawn() bool {
	return t.state == GameTileDrawn
}

// IsInHand returns true if this tile is in hand or just drawn
func (t *GameTile) IsInHand() bool {
	return t.state == GameTileDrawn || t.state == GameTileInHand
}
