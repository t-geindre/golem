package component

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image"
)

//go:generate golem component Text
type Text struct {
	Text string
	Face text.Face
	// computed
	Bounds image.Rectangle
}

func NewText(value string, face text.Face) *Text {
	return &Text{
		Text: value,
		Face: face,
	}
}
