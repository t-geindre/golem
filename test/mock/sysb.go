package mock

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
)

type SysB struct {
}

func NewSysB() *SysB {
	return &SysB{}
}

func (s *SysB) Update(e golem.Entity, _ golem.World, _ golem.Clock) {
	s.update(e)
}

func (s *SysB) Draw(e golem.Entity, _ *ebiten.Image, _ golem.World) {
	s.update(e)
}

func (s *SysB) update(e golem.Entity) {
	compa := GetCompA(e)
	if compa != nil {
		compa.Value++
	}

	compb := GetCompB(e)
	if compb != nil {
		compb.Value++
	}
}
