package system

import (
	"github.com/t-geindre/golem/pkg/golem"
	"math/rand"
	"time"
)

type SpawnFunc func(l golem.LayerID, px, py float64) golem.Entity

type Spawner struct {
	xMin, xMax, y float64
	spawn         SpawnFunc
	rate          time.Duration
	last          time.Time
	layer         golem.LayerID
}

func NewSpawner(l golem.LayerID, xMin, xMax, y float64, spawn SpawnFunc, rate time.Duration) *Spawner {
	return &Spawner{
		xMin:  xMin,
		xMax:  xMax,
		y:     y,
		spawn: spawn,
		rate:  rate,
		last:  time.Now(),
		layer: l,
	}
}

func (s *Spawner) UpdateOnce(w golem.World) {
	if time.Since(s.last) > s.rate {
		s.last = time.Now()
		e := s.spawn(s.layer, s.xMin+rand.Float64()*(s.xMax-s.xMin), s.y)
		w.AddEntity(e)
	}
}
