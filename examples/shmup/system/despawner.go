package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Despawner struct {
	wh float64
}

func NewDespawner() *Despawner {
	return &Despawner{}
}

func (d *Despawner) Update(e golem.Entity, w golem.World, _ golem.Clock) {
	life := component.GetLife(e)
	position := component.GetPosition(e)

	if life != nil && life.Current <= 0 {
		w.RemoveEntity(e)
		if life.DeathSpawn != nil && position != nil {
			w.AddEntity(life.DeathSpawn(e.GetLayer(), position.X, position.Y))
		}
		return
	}

	lifetime := component.GetLifetime(e)
	if lifetime != nil {
		if lifetime.Start.IsZero() {
			lifetime.Start = time.Now()
		}
		if time.Since(lifetime.Start) > lifetime.Life {
			w.RemoveEntity(e)
			return
		}
	}

	despawn := component.GetDespawn(e)
	if despawn == nil || position == nil {
		return
	}

	yMin, yMax := 0.0, d.wh

	sp := component.GetSprite(e)
	if sp != nil {
		sh := float64(sp.Img.Bounds().Dy())
		yMin -= sh
		yMax += sh
	}

	if despawn.Direction == component.DespawnDirectionUp && position.Y < yMin {
		w.RemoveEntity(e)
		return
	}

	if despawn.Direction == component.DespawnDirectionDown && position.Y > yMax {
		w.RemoveEntity(e)
	}
}
