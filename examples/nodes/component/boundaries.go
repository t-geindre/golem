package component

import "image"

//go:generate golem component Boundaries
type Boundaries struct {
	image.Rectangle
}

func NewBoundaries(x, y, w, h int) *Boundaries {
	return &Boundaries{Rectangle: image.Rect(x, y, w+x, h+y)}
}
