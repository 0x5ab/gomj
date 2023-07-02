package tiles

import (
	"sort"

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

func SortTilesInPlace(tiles []Tile) {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i].Id < tiles[j].Id
	})
}

func SortTiles(tiles []Tile) []Tile {
	t := make([]Tile, len(tiles))
	copy(t, tiles)
	SortTilesInPlace(t)
	return t
}
