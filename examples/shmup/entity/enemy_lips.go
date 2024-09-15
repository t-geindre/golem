package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
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

func NewEnemy(l golem.LayerID, px, py float64) golem.Entity {
	img := helper.Assets["enemy_f1"]
	w, h := float64(img.Bounds().Dx())*.7, float64(img.Bounds().Dy())*.7
	return &Enemy{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(px, py),
		Velocity: component.NewVelocity(0, 4),
		Sprite:   component.NewSprite(img),
		Despawn:  component.NewDespawn(component.DespawnDirectionDown),
		Collider: component.NewCollider(-(w / 2), -(h / 2), w, h),
		Life: component.NewLife(1, func(e golem.Entity) golem.Entity {
			pos := component.GetPosition(e)
			if pos == nil {
				panic("no position component")
			}
			return NewExplosion(e.GetLayer(), pos.X, pos.Y)
		}),
		Animation: component.NewAnimation(
			true,
			component.NewFrame(helper.Assets["enemy_f1"], time.Millisecond*500),
			component.NewFrame(helper.Assets["enemy_f2"], time.Millisecond*50),
			component.NewFrame(helper.Assets["enemy_f3"], time.Millisecond*50),
			component.NewFrame(helper.Assets["enemy_f4"], time.Millisecond*200),
			component.NewFrame(helper.Assets["enemy_f5"], time.Millisecond*50),
		),
	}
}
