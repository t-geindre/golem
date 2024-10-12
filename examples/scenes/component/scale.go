package component

//go:generate golem component Scale
type Scale struct {
	Value            float64
	OriginX, OriginY float64
}

func NewScale(v, ox, oy float64) *Scale {
	return &Scale{
		Value:   v,
		OriginX: ox,
		OriginY: oy,
	}
}
