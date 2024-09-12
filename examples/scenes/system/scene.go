package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"sync"
)

const margin = 5

type Scene struct {
	scenes  []golem.Entity
	current int
	once    sync.Once
}

func NewScene(scenes ...golem.Entity) *Scene {
	return &Scene{
		scenes: scenes,
	}
}

func (s *Scene) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	scene := component.GetScene(e)
	if scene != nil {
		ebitenutil.DebugPrintAt(screen, scene.Name, margin, margin)
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
}

func (s *Scene) switchScene(w golem.World) {
	w.RemoveEntity(s.scenes[s.current])
	s.current++
	if s.current >= len(s.scenes) {
		s.current = 0
	}
	w.AddEntity(s.scenes[s.current])
}
