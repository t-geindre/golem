package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Enemy struct {
	golem.Entity
	*component.Position
	*component.Velocity
	*component.Sprite
	*component.Despawn
	*component.Collider
	*component.Life
	*component.Animation
}

func NewEnemy(l golem.LayerID, px, py float64, frames ...component.Frame) *Enemy {
	return &Enemy{
		Entity:    golem.NewEntity(l),
		Position:  component.NewPosition(px, py),
		Velocity:  component.NewVelocity(0, 3),
		Sprite:    component.NewSprite(frames[0].Img),
		Despawn:   component.NewDespawn(component.DespawnDirectionDown),
		Collider:  component.NewCollider(-15, -15, 30, 30),
		Life:      component.NewLife(1, NewExplosion),
		Animation: component.NewAnimation(true, frames...),
	}
}
