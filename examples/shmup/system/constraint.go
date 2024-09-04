package system

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Constraint struct {
}

func NewConstraint() *Constraint {
	return &Constraint{}
}

func (c *Constraint) Update(e golem.Entity, w golem.World) {
	cs, ok := e.(component.Constraint)
	if !ok {
		return
	}
	constraint := cs.GetConstraint()

	pos, ok := e.(component.Position)
	if !ok {
		return
	}
	position := pos.GetPosition()

	if position.X < constraint.XMin {
		position.X = constraint.XMin
	}

	if position.X > constraint.XMax {
		position.X = constraint.XMax
	}

	if position.Y < constraint.YMin {
		position.Y = constraint.YMin
	}

	if position.Y > constraint.YMax {
		position.Y = constraint.YMax
	}
}
