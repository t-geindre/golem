package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Mouse struct {
	x, y     int
	lastMove time.Time
	visible  bool
}

func NewMouse() *Mouse {
	return &Mouse{
		visible: true,
	}
}

func (m *Mouse) UpdateOnce(_ golem.World, c golem.Clock) {
	x, y := ebiten.CursorPosition()
	if m.x != x || m.y != y {
		m.x, m.y = x, y
		m.lastMove = c.Now()
		if !m.visible {
			ebiten.SetCursorMode(ebiten.CursorModeVisible)
			m.visible = true
		}
		return
	}
	if c.Since(m.lastMove) > time.Second && m.visible {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		m.visible = false
	}
}
