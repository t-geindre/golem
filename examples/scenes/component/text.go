package component

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/sfnt"
)

//go:generate golem component Text
type Text struct {
	value    string
	Font     *sfnt.Font
	FontSize float64 // rem, relative to root size as defined by the renderer

	// computed
	Dirty         bool
	FontFace      text.Face
	FontAscent    float64
	Width, Height float64
}

func NewText(value string, font *sfnt.Font, size float64) *Text {
	return &Text{
		value:    value,
		Font:     font,
		FontSize: size,
	}
}

func (t *Text) SetValue(value string) {
	t.Dirty = true
	t.value = value
}

func (t *Text) GetValue() string {
	return t.value
}
