package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Animation struct {
	golem.Entity
	*component.Sprite
	*component.Animation
	*component.Position
	*component.Scale
	*component.Opacity
	*component.Boundaries
}

func NewAnimation(l golem.LayerID, fs ...component.Frame) *Animation {
	return &Animation{
		Entity:    golem.NewEntity(l),
		Sprite:    component.NewSprite(fs[0].Img),
		Animation: component.NewAnimation(fs...),
		Position:  component.NewPosition(.5, .5, .5, .5),
		Scale:     component.NewScale(1, .5, .5),
		Opacity:   component.NewOpacity(1),
		Boundaries: &component.Boundaries{
			Rectangle: fs[0].Img.Bounds(),
		},
	}
}
