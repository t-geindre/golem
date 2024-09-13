package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/squares/entity"
	"github.com/t-geindre/golem/examples/squares/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"image"
	"math/rand"
	"time"
)

type Spawner struct {
	rate  int
	last  time.Time
	layer golem.LayerID
	rect  *image.Point
	count int
	*golemutils.Panel
}

func NewSpawner(sl, dl golem.LayerID, rect *image.Point) *Spawner {
	s := &Spawner{
		rate:  1,
		layer: sl,
		rect:  rect,
	}
	s.Panel = golemutils.NewPanel(dl, s.getDisplay, golemutils.RefreshOnce)
	s.Stick = golemutils.StickTopRight
	return s
}

func (s *Spawner) UpdateOnce(w golem.World) {
	_, y := ebiten.Wheel()
	if y != 0 {
		s.rate += int(y * 5)
		if s.rate < 1 {
			s.rate = 1
		}
		s.Panel.Refresh(w)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for i := 0; i < s.rate; i++ {
			asset := helper.Assets[s.count%len(helper.Assets)]
			mw, mh := float64(s.rect.X-asset.Bounds().Dx()/2), float64(s.rect.Y-asset.Bounds().Dy())

			e := entity.NewSquare(
				s.layer,
				asset,
				rand.Float64()*mw, rand.Float64()*mh,
				(rand.Float64()*2-1)*3, (rand.Float64()*2-1)*3,
			)

			w.AddEntity(e)
			s.count++
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		target := s.rate
		for _, e := range w.GetEntities(s.layer) {
			w.RemoveEntity(e)
			target--
			if target <= 0 {
				return
			}
		}
	}

	s.Panel.UpdateOnce(w)
}

func (s *Spawner) getDisplay(w golem.World) string {
	return fmt.Sprintf("Amount: %d [MW] +/-\n[LMB]/[RMB] Add/Remove", s.rate)
}
