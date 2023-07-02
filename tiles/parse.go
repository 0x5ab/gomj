package tiles

import (
	"strconv"
	"strings"

	errors2 "errors"

	"github.com/0x5ab/gomj/errors"
)

const tileTypeStr = "mpsz"

func isTileTypeChar(c rune) bool {
	return strings.ContainsRune(tileTypeStr, c)
}

func isNumberChar(c rune) bool {
	return c >= '0' && c <= '9'
}

func ParseTile(t string) (Tile, error) {
	if len(t) != 2 {
		return Invalid, errors.ErrParseTiles
	}
	n, err := strconv.Atoi(t[0:1])
	if err != nil {
		return Invalid, errors2.Join(errors.ErrParseTiles, err)
	}
	res, err := GetTile(TileType(t[1]), n)
	if err != nil {
		return Invalid, errors2.Join(errors.ErrParseTiles, err)
	}
	return res, nil
}

func ParseTiles(s string) ([]Tile, error) {
	var tiles []Tile
	numbers := make([]int, 0, len(s))
	for _, c := range s {
		if isNumberChar(c) {
			numbers = append(numbers, int(c-'0'))
		} else if isTileTypeChar(c) {
			if len(numbers) == 0 {
				return nil, errors.ErrParseTiles
			}
			for _, n := range numbers {
				t, err := GetTile(TileType(c), n)
				if err != nil {
					return nil, errors2.Join(errors.ErrParseTiles, err)
				}
				tiles = append(tiles, t)
			}
			numbers = numbers[:0]
		} else {
			return nil, errors.ErrParseTiles
		}
	}
	if len(numbers) != 0 || len(tiles) == 0 {
		return nil, errors.ErrParseTiles
	}
	return tiles, nil
}
