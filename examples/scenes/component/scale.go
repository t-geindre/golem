package component

//go:generate golem component Scale
type Scale struct {
	Value float64
}

func NewScale(v float64) *Scale {
	return &Scale{Value: v}
}
