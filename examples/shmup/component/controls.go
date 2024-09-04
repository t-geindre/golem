package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
)

type Controls interface {
	GetControl() *ControlsImpl
}

func NewControls(up, down, left, right, fire ebiten.Key, velocity float64) *ControlsImpl {
	return &ControlsImpl{
		Up:       up,
		Down:     down,
		Left:     left,
		Right:    right,
		Fire:     fire,
		Velocity: velocity,
	}
}

type ControlsImpl struct {
	Up       ebiten.Key
	Down     ebiten.Key
	Left     ebiten.Key
	Right    ebiten.Key
	Fire     ebiten.Key
	Velocity float64
}

func (c *ControlsImpl) GetControl() *ControlsImpl {
	return c
}

func GetControls(e golem.Entity) *ControlsImpl {
	if c, ok := e.(Controls); ok {
		return c.GetControl()
	}
	return nil
}
