package component

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
