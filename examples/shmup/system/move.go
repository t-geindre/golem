package system

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Move struct {
}

func NewMove() *Move {
	return &Move{}
}

func (s *Move) Update(e golem.Entity, w golem.World) {
	pos := component.GetPosition(e)
	vel := component.GetVelocity(e)

	if pos == nil || vel == nil {
		return
	}

	pos.X += vel.X
	pos.Y += vel.Y

	collider := component.GetCollider(e)
	if collider == nil {
		return
	}

	collider.Px = pos.X + collider.ShiftX
	collider.Py = pos.Y + collider.ShiftY
}
