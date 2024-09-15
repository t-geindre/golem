package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Sparkle struct {
	golem.Entity
	*component.Position
	*component.Velocity
	*component.Sprite
	*component.Despawn
	*component.Animation
}

func NewSparkle(l golem.LayerID, px, py float64) golem.Entity {
	s := &Sparkle{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(px, py),
		Velocity: component.NewVelocity(0, 2),
		Sprite:   component.NewSprite(helper.Assets["sparkle_a"]),
		Despawn:  component.NewDespawn(component.DespawnDirectionDown),
		Animation: component.NewAnimation(
			true,
			component.NewFrame(helper.Assets["sparkle_f1"], time.Millisecond*200),
			component.NewFrame(helper.Assets["sparkle_f2"], time.Millisecond*200),
			component.NewFrame(helper.Assets["sparkle_f3"], time.Millisecond*200),
			component.NewFrame(helper.Assets["sparkle_f4"], time.Millisecond*200),
			component.NewFrame(helper.Assets["sparkle_f5"], time.Millisecond*200),
			component.NewFrame(helper.Assets["sparkle_f4"], time.Millisecond*200),
			component.NewFrame(helper.Assets["sparkle_f3"], time.Millisecond*200),
			component.NewFrame(helper.Assets["sparkle_f2"], time.Millisecond*200),
		),
	}
	return s
}
