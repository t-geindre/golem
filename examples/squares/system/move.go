package system

import (
	"github.com/t-geindre/golem/examples/squares/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Move struct {
}

func NewMove() *Move {
	return &Move{}
}

func (s *Move) Update(e golem.Entity, _ golem.World, _ golem.Clock) {
	pos := component.GetPosition(e)
	vel := component.GetVelocity(e)

	if pos == nil || vel == nil {
		return
	}

	pos.X += vel.X
	pos.Y += vel.Y
}
