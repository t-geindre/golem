package system

import (
	"github.com/t-geindre/golem/examples/nodes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Move struct {
}

func NewMove() *Move {
	return &Move{}
}

func (m *Move) Update(e golem.Entity, _ golem.World, _ golem.Clock) {
	vel := component.GetVelocity(e)
	if vel == nil {
		return
	}

	bds := component.GetBoundaries(e)
	if bds == nil {
		return
	}

	bds.Rectangle = bds.Add(vel.Point)
}
