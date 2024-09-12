package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

//go:generate golem component Controls
type Controls struct {
	Up       ebiten.Key
	Down     ebiten.Key
	Left     ebiten.Key
	Right    ebiten.Key
	Fire     ebiten.Key
	Velocity float64
}

func NewControls(up, down, left, right, fire ebiten.Key, velocity float64) *Controls {
	return &Controls{
		Up:       up,
		Down:     down,
		Left:     left,
		Right:    right,
		Fire:     fire,
		Velocity: velocity,
	}
}
