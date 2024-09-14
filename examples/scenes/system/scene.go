package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"sync"
)

const margin = 5

type Scene struct {
	scenes  []golem.Entity
	current int
	once    sync.Once
	*golemutils.Panel
}

func NewScene(l golem.LayerID, scenes ...golem.Entity) *Scene {
	s := &Scene{
		scenes: scenes,
	}
	s.Panel = golemutils.NewPanel(l, s.getCurrentScene, golemutils.RefreshOnce)
	s.Panel.Stick = golemutils.StickBottomRight
	return s
}

func (s *Scene) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	scene := component.GetScene(e)
	if scene != nil {
		scene.Draw(screen)
	}
}

func (s *Scene) Update(e golem.Entity, w golem.World) {
	scene := component.GetScene(e)
	if scene != nil {
		scene.Update()
	}
}

func (s *Scene) UpdateOnce(w golem.World) {
	if len(s.scenes) == 0 {
		return
	}

	s.once.Do(func() {
		s.switchScene(w)
	})

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.switchScene(w)
	}

	s.Panel.UpdateOnce(w)
}

func (s *Scene) getCurrentScene(w golem.World) string {
	c := s.scenes[s.current]
	sc := component.GetScene(c)
	if s != nil {
		return fmt.Sprintf("[LMB] Switch scene\nCurrent scene : %s", sc.Name)
	}
	return ""
}

func (s *Scene) switchScene(w golem.World) {
	w.RemoveEntity(s.scenes[s.current])
	s.current++
	if s.current >= len(s.scenes) {
		s.current = 0
	}
	w.AddEntity(s.scenes[s.current])
	s.Panel.Refresh(w)
}
