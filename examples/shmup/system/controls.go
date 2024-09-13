package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Controls struct {
}

func NewControls() *Controls {
	return &Controls{}
}

func (c *Controls) Update(e golem.Entity, w golem.World) {
	controls := component.GetControls(e)
	if controls == nil {
		return
	}

	velocity := component.GetVelocity(e)
	if velocity != nil {
		if ebiten.IsKeyPressed(controls.Up) {
			velocity.Y = -controls.Velocity
		} else if ebiten.IsKeyPressed(controls.Down) {
			velocity.Y = controls.Velocity
		} else {
			velocity.Y = 0
		}

		if ebiten.IsKeyPressed(controls.Left) {
			velocity.X = -controls.Velocity
		} else if ebiten.IsKeyPressed(controls.Right) {
			velocity.X = controls.Velocity
		} else {
			velocity.X = 0
		}

		if velocity.X != 0 && velocity.Y != 0 {
			velocity.X *= 0.70710678118
			velocity.Y *= 0.70710678118
		}
	}

	shoot := component.GetShoot(e)
	if shoot == nil {
		return
	}

	shoot.Shooting = false
	if ebiten.IsKeyPressed(controls.Fire) {
		shoot.Shooting = true
	}
}
