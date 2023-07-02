package errors

import "errors"

var (
	ErrInvalidTile = errors.New("invalid tile")
	ErrParseTiles  = errors.New("parse tiles error")
)
