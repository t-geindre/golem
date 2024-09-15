package component

//go:generate golem component Constraint
type Constraint struct {
	X, Y float64
	W, H float64
}

func NewConstraint(x, y, w, h float64) *Constraint {
	return &Constraint{x, y, w, h}
}
