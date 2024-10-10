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

func (s *Shoot) Update(e golem.Entity, w golem.World, _ golem.Clock) {
	shoot := component.GetShoot(e)
	if shoot == nil || !shoot.Shooting {
		return
	}

	if time.Since(shoot.Last) < shoot.Rate {
		return
	}

	shoot.Last = time.Now()
	bullet := shoot.Spawn(shoot.Layer)

	w.AddEntity(bullet)

	bPos := component.GetPosition(bullet)
	srcPos := component.GetPosition(e)

	if bPos == nil || srcPos == nil {
		return
	}

	bPos.X = srcPos.X + shoot.AtX
	bPos.Y = srcPos.Y + shoot.AtY
}
