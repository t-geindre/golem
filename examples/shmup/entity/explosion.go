package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Explosion struct {
	golem.Entity
	*component.Position
	*component.Sprite
	*component.Animation
	*component.Lifetime
}

func NewExplosion(l golem.LayerID, px, py float64) golem.Entity {
	return &Explosion{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(px, py),
		Sprite:   component.NewSprite(helper.Assets["explosion_f1"]),
		Animation: component.NewAnimation(
			false,
			component.NewFrame(helper.Assets["explosion_f1"], time.Millisecond*40),
			component.NewFrame(helper.Assets["explosion_f2"], time.Millisecond*40),
			component.NewFrame(helper.Assets["explosion_f3"], time.Millisecond*40),
			component.NewFrame(helper.Assets["explosion_f4"], time.Millisecond*40),
			component.NewFrame(helper.Assets["explosion_f5"], time.Millisecond*40),
		),
		Lifetime: component.NewLifetime(time.Millisecond * 5 * 40),
	}
}
