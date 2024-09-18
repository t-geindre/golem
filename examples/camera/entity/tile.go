package entity

import (
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type Tile struct {
	golem.Entity
	*component.Animation
	*component.Position
	*component.Sprite
}

func NewTile(l golem.LayerID, p image.Point, fs ...component.Frame) golem.Entity {
	t := &Tile{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(p),
		Sprite:   component.NewSprite(fs[0].Img),
	}

	if len(fs) > 1 {
		t.Animation = component.NewAnimation(true, fs...)
	}

	return t
}
