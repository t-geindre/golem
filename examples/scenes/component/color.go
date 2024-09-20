package component

import "image/color"

//go:generate golem component Color
type Color struct {
	Value color.RGBA
}

func NewColor(col color.RGBA) *Color {
	return &Color{
		Value: col,
	}
}
