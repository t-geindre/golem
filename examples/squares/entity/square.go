package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/squares/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Square struct {
	golem.Entity
	component.Position
	component.Velocity
	component.Sprite
}

func NewSquare(l golem.LayerID, img *ebiten.Image, px, py, vx, vy float64) golem.Entity {
	return &Square{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(px, py),
		Velocity: component.NewVelocity(vx, vy),
		Sprite:   component.NewSprite(img),
	}
}
