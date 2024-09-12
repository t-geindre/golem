package component

//go:generate golem component Constraint
type Constraint struct {
	XMin, XMax float64
	YMin, YMax float64
}

func NewConstraint(xMin, xMax, yMin, yMax float64) *Constraint {
	return &Constraint{xMin, xMax, yMin, yMax}
}
