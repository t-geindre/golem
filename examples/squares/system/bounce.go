package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/squares/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Bounce struct {
}

func NewBounce() *Bounce {
	return &Bounce{}
}

func (s *Bounce) Update(e golem.Entity, w golem.World) {
	pos := component.GetPosition(e)
	vel := component.GetVelocity(e)

	if pos == nil || vel == nil {
		return
	}

	mx, my := 0, 0
	sp := component.GetSprite(e)
	if sp != nil {
		spr := sp.GetSprite()
		mx, my = spr.Img.Bounds().Dx()/2, spr.Img.Bounds().Dy()/2
	}
	ww, wh := ebiten.WindowSize()

	if pos.X < float64(mx) {
		pos.X = float64(mx)
		vel.X = -vel.X
	}

	if pos.Y < float64(my) {
		pos.Y = float64(my)
		vel.Y = -vel.Y
	}

	if pos.X > float64(ww-mx) {
		pos.X = float64(ww - mx)
		vel.X = -vel.X
	}

	if pos.Y > float64(wh-my) {
		pos.Y = float64(wh - my)
		vel.Y = -vel.Y
	}
}
