package component

import "image"

//go:generate golem component Position
type Position struct {
	image.Point
}

func NewPosition(p image.Point) *Position {
	return &Position{Point: p}
}
