package entity

import (
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Tile struct {
	golem.Entity
	*component.Animation
	*component.Position
	*component.Sprite
}

func NewTile(l golem.LayerID, x, y int, fs ...component.Frame) golem.Entity {
	return &Tile{
		Entity:    golem.NewEntity(l),
		Position:  component.NewPosition(x, y),
		Sprite:    component.NewSprite(fs[0].Img),
		Animation: component.NewAnimation(true, fs...),
	}
}
