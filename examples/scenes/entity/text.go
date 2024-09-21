package entity

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"image/color"
)

type Text struct {
	golem.Entity
	*component.Text
	*component.Position
	*component.Color
}

func NewText(l golem.LayerID, text string, x, y, ox, oy float64, face text.Face, col color.RGBA) *Text {
	return &Text{
		Entity:   golem.NewEntity(l),
		Text:     component.NewText(text, face),
		Position: component.NewPosition(x, y, ox, oy),
		Color:    component.NewColor(col),
	}
}
