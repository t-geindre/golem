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
}

func NewGopher(l golem.LayerID, fs ...component.Frame) *Gopher {
	return &Gopher{
		Entity:    golem.NewEntity(l),
		Position:  component.NewPosition(0.5, 0.5),
		Sprite:    component.NewSprite(fs[0].Img),
		Animation: component.NewAnimation(fs...),
	}
}
