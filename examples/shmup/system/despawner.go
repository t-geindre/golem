package system

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
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
	if life != nil && life.Current <= 0 {
		w.RemoveEntity(e)
	}

	despawn := component.GetDespawn(e)
	position := component.GetPosition(e)

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
