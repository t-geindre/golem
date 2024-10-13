package component

import "image"

//go:generate golem component Velocity
type Velocity struct {
	image.Point
}

func NewVelocity(x, y int) *Velocity {
	return &Velocity{Point: image.Pt(x, y)}
}
