package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/squares/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Bounce struct {
	bx, by float64
}

func NewBounce() *Bounce {
	return &Bounce{}
}

func (s *Bounce) UpdateOnce(_ golem.World, _ golem.Clock) {
	ww, wh := ebiten.WindowSize()
	s.bx, s.by = float64(ww), float64(wh)
}

func (s *Bounce) Update(e golem.Entity, _ golem.World, _ golem.Clock) {
	pos := component.GetPosition(e)
	vel := component.GetVelocity(e)

	if pos == nil || vel == nil {
		return
	}

	mx, my := 0, 0
	sp := component.GetSprite(e)
	if sp != nil {
		mx, my = sp.Img.Bounds().Dx()/2, sp.Img.Bounds().Dy()/2
	}

	if pos.X < float64(mx) {
		pos.X = float64(mx)
		vel.X = -vel.X
	}

	if pos.Y < float64(my) {
		pos.Y = float64(my)
		vel.Y = -vel.Y
	}

	if pos.X > s.bx-float64(mx) {
		pos.X = s.bx - float64(mx)
		vel.X = -vel.X
	}

	if pos.Y > s.by-float64(my) {
		pos.Y = s.by - float64(my)
		vel.Y = -vel.Y
	}
}
