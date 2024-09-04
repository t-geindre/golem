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
	ctrl, ok := e.(component.Controls)
	if !ok {
		return
	}
	controls := ctrl.GetControl()

	vel, ok := e.(component.Velocity)
	if !ok {
		return
	}
	velocity := vel.GetVelocity()

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

	sh, ok := e.(component.Shoot)
	if !ok {
		return
	}
	shoot := sh.GetShoot()

	shoot.Shooting = false
	if ebiten.IsKeyPressed(controls.Fire) {
		shoot.Shooting = true
	}
}
