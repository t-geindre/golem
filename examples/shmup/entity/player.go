package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Player struct {
	golem.Entity
	*component.Position
	*component.Velocity
	*component.Sprite
	*component.Controls
	*component.Shoot
	*component.Constraint
	*component.Collider
	*component.Life
}

func NewPlayer(l, bl golem.LayerID, px, py, cxMin, cxMax, cyMin, cyMax float64) golem.Entity {
	img := helper.Assets["player"]
	w, h := float64(img.Bounds().Dx())*.7, float64(img.Bounds().Dy())*.8

	return &Player{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(px, py),
		Velocity: component.NewVelocity(0, 0),
		Sprite:   component.NewSprite(img),
		Controls: component.NewControls(
			ebiten.KeyUp,
			ebiten.KeyDown,
			ebiten.KeyLeft,
			ebiten.KeyRight,
			ebiten.KeySpace,
			5,
		),
		Shoot:      component.NewShoot(time.Millisecond*150, 0, -h*.7, NewBullet, bl),
		Constraint: component.NewConstraint(cxMin, cxMax, cyMin, cyMax),
		Collider:   component.NewCollider(-(w / 2), -(h / 2), w, h),
		Life:       component.NewLife(5),
	}
}
