package component

type ScreenScale struct {
	Value float64
}

//go:generate golem component ScreenScale
func NewScreenScale() *ScreenScale {
	return &ScreenScale{
		Value: 1,
	}
}
