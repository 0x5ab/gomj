package tiles

import (
	mapset "github.com/deckarep/golang-set/v2"
)

func HasDuplicateTile(tiles []Tile) bool {
	tileSet := mapset.NewSet[int]()
	for _, tile := range tiles {
		tileSet.Add(tile.Number)
	}
	return tileSet.Cardinality() != len(tiles)
}

func CountDuplicateTiles(tiles []Tile) int {
	tileSet := mapset.NewSet[int]()
	for _, tile := range tiles {
		tileSet.Add(tile.Number)
	}
	return len(tiles) - tileSet.Cardinality()
}
