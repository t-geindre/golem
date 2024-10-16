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
	*component.Scale
	*component.Opacity
	*component.Boundaries
}

func NewText(l golem.LayerID, text string) *Text {
	return &Text{
		Entity:     golem.NewEntity(l),
		Text:       component.NewText(text),
		Position:   component.NewPosition(.5, .5, .5, .5),
		Color:      component.NewColor(colornames.Black),
		Scale:      component.NewScale(1, .5, .5),
		Opacity:    component.NewOpacity(1),
		Boundaries: component.NewBoundaries(0, 0, 0, 0),
	}
}
