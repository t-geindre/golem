package component

//go:generate golem position.go
type Position struct {
	X, Y float64
}

func NewPosition(x, y float64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}
