package mock

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
)

type SysC struct {
}

func NewSysC() *SysC {
	return &SysC{}
}

func (s *SysC) Update(e golem.Entity, _ golem.World) {
	s.update(e)
}

func (s *SysC) Draw(e golem.Entity, _ *ebiten.Image, _ golem.World) {
	s.update(e)
}

func (s *SysC) update(e golem.Entity) {
	compa := GetCompA(e)
	if compa != nil {
		compa.Value++
	}

	compb := GetCompB(e)
	if compb != nil {
		compb.Value++
	}

	compc := GetCompC(e)
	if compc != nil {
		compc.Value++
	}
}
