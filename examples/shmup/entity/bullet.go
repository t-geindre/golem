package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Bullet struct {
	golem.Entity
	*component.Sprite
	*component.Position
	*component.Velocity
	*component.Despawn
	*component.Collider
	*component.Life
	*component.Animation
}

func NewBullet(l golem.LayerID) golem.Entity {
	return &Bullet{
		Entity:   golem.NewEntity(l),
		Sprite:   component.NewSprite(helper.Assets["bullet_f1"]),
		Position: component.NewPosition(0, 0),
		Velocity: component.NewVelocity(0, -8),
		Despawn:  component.NewDespawn(component.DespawnDirectionUp),
		Collider: component.NewCollider(-4, -3, 8, 6),
		Life:     component.NewLife(1),
		Animation: component.NewAnimation(
			component.NewFrame(helper.Assets["bullet_f1"], time.Millisecond*100),
			component.NewFrame(helper.Assets["bullet_f2"], time.Millisecond*100),
		),
	}
}
