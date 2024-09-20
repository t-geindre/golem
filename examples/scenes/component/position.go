package component

//go:generate golem component Position
type Position struct {
	// Relative position
	RelX, RelY float64
}

func NewPosition(x, y float64) *Position {
	return &Position{
		RelX: x,
		RelY: y,
	}
}
