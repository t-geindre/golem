package component

//go:generate golem collider.go
type Collider struct {
	Px, Py, ShiftX, ShiftY, Width, Height float64
}

func NewCollider(sx, sy, w, h float64) *Collider {
	return &Collider{
		ShiftX: sx,
		ShiftY: sy,
		Width:  w,
		Height: h,
	}
}
