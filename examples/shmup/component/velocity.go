package component

//go:generate golem velocity.go
type Velocity struct {
	X, Y float64
}

func NewVelocity(x, y float64) *Velocity {
	return &Velocity{
		X: x,
		Y: y,
	}
}
