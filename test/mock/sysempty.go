package mock

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
)

type SysEmpty struct {
}

func NewSysEmpty() *SysEmpty {
	return &SysEmpty{}
}

func (s *SysEmpty) Update(e golem.Entity, _ golem.World) {
}

func (s *SysEmpty) Draw(e golem.Entity, _ *ebiten.Image, _ golem.World) {
}
