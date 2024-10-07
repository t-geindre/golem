package component

//go:generate golem component Rotation
type Rotation struct {
	Angle           float64
	RotationOriginX float64
	RotationOriginY float64
}

func NewRotation(a float64) *Rotation {
	return &Rotation{
		Angle: a,
	}
}
