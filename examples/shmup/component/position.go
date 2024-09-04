package component

import "github.com/t-geindre/golem/pkg/golem"

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

func GetPosition(e golem.Entity) *PositionImpl {
	if p, ok := e.(Position); ok {
		return p.GetPosition()
	}
	return nil
}
