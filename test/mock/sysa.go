package mock

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
)

type SysA struct {
}

func NewSysA() *SysA {
	return &SysA{}
}

func (s *SysA) Update(e golem.Entity, _ golem.World, _ golem.Clock) {
	s.update(e)
}

func (s *SysA) Draw(e golem.Entity, _ *ebiten.Image, _ golem.World) {
	s.update(e)
}

func (s *SysA) update(e golem.Entity) {
	compa := GetCompA(e)
	if compa == nil {
		return
	}
	compa.Value++
}
