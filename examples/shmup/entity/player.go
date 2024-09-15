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
	*component.Animation
}

func NewPlayer(l, bl golem.LayerID) golem.Entity {
	ww, wh := ebiten.WindowSize()
	sprite := helper.Assets["player_f1"]

	px := float64(ww)/2 - float64(sprite.Bounds().Dx())/2
	py := float64(wh) - float64(sprite.Bounds().Dy())

	return &Player{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(px, py),
		Velocity: component.NewVelocity(0, 0),
		Sprite:   component.NewSprite(sprite),
		Controls: component.NewControls(
			ebiten.KeyUp,
			ebiten.KeyDown,
			ebiten.KeyLeft,
			ebiten.KeyRight,
			ebiten.KeySpace,
			5,
		),
		Shoot:      component.NewShoot(time.Millisecond*150, 0, -32, NewBullet, bl),
		Constraint: component.NewConstraint(-16, -32, 32, 32),
		Collider:   component.NewCollider(-13, -29, 26, 26),
		Life:       component.NewLife(5, NewExplosion),
		Animation: component.NewAnimation(
			true,
			component.NewFrame(helper.Assets["player_f1"], time.Millisecond*50),
			component.NewFrame(helper.Assets["player_f2"], time.Millisecond*50),
		),
	}
}
