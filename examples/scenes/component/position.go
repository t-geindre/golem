package component

//go:generate golem component Position
type Position struct {
	// Relative position
	RelX, RelY float64
	// On item origin
	OrigX, OrigY float64
}

func NewPosition(x, y, ox, oy float64) *Position {
	return &Position{
		RelX:  x,
		RelY:  y,
		OrigX: ox,
		OrigY: oy,
	}
}
