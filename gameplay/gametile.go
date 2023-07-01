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
	GameTileAsDora
	GameTileAsUraDora
)

type GameTile struct {
	IsLingShang bool // is this tile ling shang
	IsLast      bool // is this last drawable tile
	Tile        tiles.Tile
	Player      *Player       // who played or is owning this tile, nil if it's a tile in the wall
	state       GameTileState // state of this tile
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

func (t *GameTile) IsInHand() bool {
	return t.state == GameTileDrawn || t.state == GameTileInHand
}
