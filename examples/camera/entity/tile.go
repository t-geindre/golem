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

func NewTile(l golem.LayerID, x, y int, frames ...component.Frame) golem.Entity {
	return &Tile{
		Entity:    golem.NewEntity(l),
		Animation: component.NewAnimation(true, frames...),
		Position:  component.NewPosition(x, y),
		Sprite:    component.NewSprite(frames[0].Img),
	}
}
