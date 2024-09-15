package system

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Despawner struct {
	yMin, yMax float64
	margin     float64
}

func NewDespawner(yMin, yMax float64, margin float64) *Despawner {
	return &Despawner{
		yMin:   yMin,
		yMax:   yMax,
		margin: margin,
	}
}

func (d *Despawner) Update(e golem.Entity, w golem.World) {
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

	if despawn.Direction == component.DespawnDirectionUp && position.Y < d.yMin-d.margin {
		w.RemoveEntity(e)
		return
	}

	if despawn.Direction == component.DespawnDirectionDown && position.Y > d.yMax+d.margin {
		w.RemoveEntity(e)
	}
}
