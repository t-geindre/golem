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

func (s *Move) Update(e golem.Entity, w golem.World) {
	if pos, vel, ok := s.checkEntity(e); ok {
		pos.X += vel.X
		pos.Y += vel.Y
	}
}

func (s *Move) checkEntity(e golem.Entity) (*component.PositionImpl, *component.VelocityImpl, bool) {
	pos, ok := e.(component.Position)
	if !ok {
		return nil, nil, false
	}

	vel, ok := e.(component.Velocity)
	if !ok {
		return nil, nil, false
	}

	return pos.GetPosition(), vel.GetVelocity(), true
}
