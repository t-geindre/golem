package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
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
}

func NewEnemy(l golem.LayerID, px, py float64) golem.Entity {
	return &Enemy{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(px, py),
		Velocity: component.NewVelocity(0, 5),
		Sprite:   component.NewSprite(helper.Assets["enemy"]),
		Despawn:  component.NewDespawn(component.DespawnDirectionDown),
		Collider: component.NewCollider(-10, -10, 20, 20),
		Life:     component.NewLife(1),
	}
}
