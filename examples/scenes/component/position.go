package component

//go:generate golem component Position
type Position struct {
	X, Y float64
}

func NewPosition(x, y float64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}
