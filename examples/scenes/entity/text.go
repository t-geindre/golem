package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"golang.org/x/image/colornames"
)

type Text struct {
	golem.Entity
	*component.Text
	*component.Position
	*component.Color
}

func NewText(l golem.LayerID, text string) *Text {
	return &Text{
		Entity:   golem.NewEntity(l),
		Text:     component.NewText(text),
		Position: component.NewPosition(.5, .5, .5, .5),
		Color:    component.NewColor(colornames.Black),
	}
}
