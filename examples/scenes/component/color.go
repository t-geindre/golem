package component

import "image/color"

//go:generate golem component Color
type Color struct {
	Value color.Color
}

func NewColor(col color.Color) *Color {
	return &Color{
		Value: col,
	}
}
