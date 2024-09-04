package system

import (
	"github.com/t-geindre/golem/examples/squares/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Bounce struct {
	bx, by float64
}

func NewBounce(bx, by float64) *Bounce {
	return &Bounce{
		bx: bx,
		by: by,
	}
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
