package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"golang.org/x/image/font/sfnt"
	"image/color"
)

type Text struct {
	golem.Entity
	*component.Text
	*component.Position
	*component.Color
}

func NewText(l golem.LayerID, text string, x, y, ox, oy float64, font *sfnt.Font, size float64, col color.RGBA) *Text {
	return &Text{
		Entity:   golem.NewEntity(l),
		Text:     component.NewText(text, font, size),
		Position: component.NewPosition(x, y, ox, oy),
		Color:    component.NewColor(col),
	}
}
