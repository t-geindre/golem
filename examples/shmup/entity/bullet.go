package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
)

type Bullet struct {
	golem.Entity
	component.Sprite
	component.Position
	component.Velocity
	component.Despawn
	component.Collider
	component.Life
}

func NewBullet(l golem.LayerID) golem.Entity {
	return &Bullet{
		Entity:   golem.NewEntity(l),
		Sprite:   component.NewSprite(helper.Assets["bullet"]),
		Position: component.NewPosition(0, 0),
		Velocity: component.NewVelocity(0, -13),
		Despawn:  component.NewDespawn(component.DespawnDirectionUp),
		Collider: component.NewCollider(-3, -3, 6, 6),
		Life:     component.NewLife(1),
	}
}
