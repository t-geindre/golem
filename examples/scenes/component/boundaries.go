package component

import "image"

//go:generate golem component Boundaries
type Boundaries struct {
	image.Rectangle
	StickScreen bool
}

func NewBoundaries(x0, y0, x1, y1 int) *Boundaries {
	return &Boundaries{
		Rectangle: image.Rect(x0, y0, x1, y1),
	}
}

func NewBoundariesStickScreen() *Boundaries {
	return &Boundaries{
		StickScreen: true,
	}
}
