package component

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/sfnt"
	"image"
)

//go:generate golem component Text
type Text struct {
	Text string
	Font *sfnt.Font
	Size float64

	// computed
	Face       text.Face
	Bounds     image.Rectangle
	LineHeight float64
	Scale      float64
}

func NewText(value string) *Text {
	return &Text{
		Text: value,
	}
}
