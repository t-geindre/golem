package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Move struct {
}

func NewMove() *Move {
	return &Move{}
}

func (s *Move) Update(e golem.Entity, _ golem.World, c golem.Clock) {
	pos := component.GetPosition(e)
	vel := component.GetVelocity(e)

	if pos == nil || vel == nil {
		return
	}

	pos.X += vel.X
	pos.Y += vel.Y

	cs := component.GetConstraint(e)
	if cs != nil {
		ww, wh := ebiten.WindowSize()
		maxX, maxY := float64(ww)-cs.W-cs.X, float64(wh)-cs.H-cs.Y
		minX, minY := -cs.X, -cs.Y

		if pos.X < minX {
			pos.X = minX
		}

		if pos.X > maxX {
			pos.X = maxX
		}

		if pos.Y < minY {
			pos.Y = minY
		}

		if pos.Y > maxY {
			pos.Y = maxY
		}
	}

	collider := component.GetCollider(e)
	if collider == nil {
		return
	}

	collider.Px = pos.X + collider.ShiftX
	collider.Py = pos.Y + collider.ShiftY
}
