package component

//go:generate golem component Rotation
type Rotation struct {
	Angle float64
}

func NewRotation(angle float64) *Rotation {
	return &Rotation{
		Angle: angle,
	}
}
