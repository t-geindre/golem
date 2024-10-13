package component

import (
	"image/color"
)

//go:generate golem component Color
type Color struct {
	Background color.Color
	Borders    color.Color
}

func NewColor(bg, bd color.Color) *Color {
	return &Color{
		Background: bg,
		Borders:    bd,
	}
}
