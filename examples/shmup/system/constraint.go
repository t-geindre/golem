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
	constraint := component.GetConstraint(e)
	position := component.GetPosition(e)

	if constraint == nil || position == nil {
		return
	}

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
