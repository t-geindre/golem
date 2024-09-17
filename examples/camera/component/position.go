package component

import "image"

//go:generate golem component Position
type Position struct {
	image.Point
}

func NewPosition(x, y int) *Position {
	return &Position{
		Point: image.Point{
			X: x,
			Y: y,
		},
	}
}
