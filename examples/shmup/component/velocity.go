package component

import "github.com/t-geindre/golem/pkg/golem"

type Velocity interface {
	GetVelocity() *VelocityImpl
}

type VelocityImpl struct {
	X, Y float64
}

func NewVelocity(x, y float64) Velocity {
	return &VelocityImpl{
		X: x,
		Y: y,
	}
}

func (v *VelocityImpl) GetVelocity() *VelocityImpl {
	return v
}

func GetVelocity(e golem.Entity) *VelocityImpl {
	if v, ok := e.(Velocity); ok {
		return v.GetVelocity()
	}
	return nil
}
