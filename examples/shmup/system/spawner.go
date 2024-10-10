package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
	"math/rand"
	"time"
)

type SpawnFunc func(l golem.LayerID) golem.Entity

type Spawner struct {
	spawns []SpawnFunc
	rate   time.Duration
	last   time.Time
	layer  golem.LayerID
}

func NewSpawner(l golem.LayerID, rate time.Duration, spawns ...SpawnFunc) *Spawner {
	return &Spawner{
		spawns: spawns,
		rate:   rate,
		last:   time.Now(),
		layer:  l,
	}
}

func (s *Spawner) UpdateOnce(w golem.World, _ golem.Clock) {
	if time.Since(s.last) > s.rate {
		s.last = time.Now()

		e := s.spawns[rand.Intn(len(s.spawns))](s.layer)

		xMax, _ := ebiten.WindowSize()
		xMin := 0

		sp := component.GetSprite(e)
		pos := component.GetPosition(e)

		if sp != nil && pos != nil {
			bd := sp.Img.Bounds()
			xMin += bd.Dx() / 2
			xMax -= bd.Dx() / 2
			pos.X = float64(rand.Intn(xMax-xMin) + xMin)
		}

		w.AddEntity(e)
	}
}
