package component

import "github.com/t-geindre/golem/pkg/golem"

type Constraint interface {
	GetConstraint() *ConstraintImpl
}

type ConstraintImpl struct {
	XMin, XMax float64
	YMin, YMax float64
}

func NewConstraint(xMin, xMax, yMin, yMax float64) *ConstraintImpl {
	return &ConstraintImpl{xMin, xMax, yMin, yMax}
}

func (c *ConstraintImpl) GetConstraint() *ConstraintImpl {
	return c
}

func GetConstraint(e golem.Entity) *ConstraintImpl {
	if c, ok := e.(Constraint); ok {
		return c.GetConstraint()
	}
	return nil
}
