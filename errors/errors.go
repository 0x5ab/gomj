package errors

import "errors"

var (
	ErrInvalidTile     = errors.New("invalid tile")
	ErrInvalidTileType = errors.New("invalid tile type")
	ErrParseTiles      = errors.New("parse tiles error")
)
