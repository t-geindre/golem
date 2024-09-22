package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Image struct {
	golem.Entity
	*component.Sprite
	*component.Position
	*component.Scale
	*component.Opacity
}

func NewImage(l golem.LayerID, img *ebiten.Image) *Image {
	return &Image{
		Entity:   golem.NewEntity(l),
		Sprite:   component.NewSprite(img),
		Position: component.NewPosition(.5, .5, .5, .5),
		Scale:    component.NewScale(1, .5, .5),
		Opacity:  component.NewOpacity(1),
	}
}
