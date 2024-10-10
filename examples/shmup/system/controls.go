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

func (c *Controls) Update(e golem.Entity, _ golem.World, _ golem.Clock) {
	controls := component.GetControls(e)
	if controls == nil {
		return
	}

	vel := component.GetVelocity(e)
	if vel != nil {
		if ebiten.IsKeyPressed(controls.Up) {
			vel.Y = -controls.Velocity
		} else if ebiten.IsKeyPressed(controls.Down) {
			vel.Y = controls.Velocity
		} else {
			vel.Y = 0
		}

		if ebiten.IsKeyPressed(controls.Left) {
			vel.X = -controls.Velocity
		} else if ebiten.IsKeyPressed(controls.Right) {
			vel.X = controls.Velocity
		} else {
			vel.X = 0
		}

		if vel.X != 0 && vel.Y != 0 {
			vel.X *= 0.70710678118
			vel.Y *= 0.70710678118
		}
	}

	animSet := component.GetAnimationSet(e)
	if animSet != nil {
		if vel.X > 0 {
			animSet.Next = "right"
		} else if vel.X < 0 {
			animSet.Next = "left"
		} else {
			animSet.Next = "idle"
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
