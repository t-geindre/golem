package system

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Shoot struct {
}

func NewShoot() *Shoot {
	return &Shoot{}
}

func (s *Shoot) Update(e golem.Entity, w golem.World) {
	sh, ok := e.(component.Shoot)
	if !ok {
		return
	}
	shoot := sh.GetShoot()

	if !shoot.Shooting {
		return
	}

	if time.Since(shoot.Last) < shoot.Rate {
		return
	}

	shoot.Last = time.Now()
	bullet := shoot.Spawn(shoot.Layer)

	w.AddEntity(bullet)

	p, ok := bullet.(component.Position)
	if !ok {
		return
	}
	bPos := p.GetPosition()

	p, ok = e.(component.Position)
	if !ok {
		return
	}
	srcPos := p.GetPosition()

	bPos.X = srcPos.X + shoot.AtX
	bPos.Y = srcPos.Y + shoot.AtY
}
