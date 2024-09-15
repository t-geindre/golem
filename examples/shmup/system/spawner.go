package system

import (
	"github.com/t-geindre/golem/pkg/golem"
	"math/rand"
	"time"
)

type SpawnFunc func(l golem.LayerID, px, py float64) golem.Entity

type Spawner struct {
	xMin, xMax, y float64
	spawns        []SpawnFunc
	rate          time.Duration
	last          time.Time
	layer         golem.LayerID
}

func NewSpawner(l golem.LayerID, xMin, xMax, y float64, rate time.Duration, spawns ...SpawnFunc) *Spawner {
	return &Spawner{
		xMin:   xMin,
		xMax:   xMax,
		y:      y,
		spawns: spawns,
		rate:   rate,
		last:   time.Now(),
		layer:  l,
	}
}

func (s *Spawner) UpdateOnce(w golem.World) {
	if time.Since(s.last) > s.rate {
		s.last = time.Now()
		e := s.spawns[rand.Intn(len(s.spawns))](s.layer, s.xMin+rand.Float64()*(s.xMax-s.xMin), s.y)
		w.AddEntity(e)
	}
}
