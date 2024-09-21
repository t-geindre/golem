package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Gopher struct {
	golem.Entity
	*component.Position
	*component.Sprite
	*component.Animation
	*component.Scale
}

func NewGopher(l golem.LayerID, fs ...component.Frame) *Gopher {
	return &Gopher{
		Entity:    golem.NewEntity(l),
		Position:  component.NewPosition(.5, .5, .5, .5),
		Sprite:    component.NewSprite(fs[0].Img),
		Animation: component.NewAnimation(fs...),
		Scale:     component.NewScale(2, .5, .5),
	}
}
