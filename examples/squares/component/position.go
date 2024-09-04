package component

type Position interface {
	GetPosition() *PositionImpl
}

type PositionImpl struct {
	X, Y float64
}

func NewPosition(x, y float64) Position {
	return &PositionImpl{
		X: x,
		Y: y,
	}
}

func (p *PositionImpl) GetPosition() *PositionImpl {
	return p
}
