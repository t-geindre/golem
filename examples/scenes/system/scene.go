package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"sync"
	"time"
)

type SceneBuilder func() golem.Entity

type Scene struct {
	scenes        []golem.Entity
	idx           int
	current, next golem.Entity
	transitioning bool
	once          sync.Once
	lastSwitch    time.Time
	*golemutils.Panel
}

func NewScene(l golem.LayerID, scenes ...golem.Entity) *Scene {
	s := &Scene{
		scenes:        scenes,
		transitioning: true, // Disable controls by default
	}

	s.Panel = golemutils.NewPanel(l, s.getPanelInfos, golemutils.RefreshOnce)
	s.Panel.Stick = golemutils.StickBottomRight

	return s
}

func (s *Scene) Update(e golem.Entity, w golem.World) {
	tr := component.GetTransition(e)
	s.transitioning = false
	if tr != nil && tr.Transitioning {
		s.transitioning = true
		isCurrent := e == s.current
		isNext := e == s.next
		v := float64(time.Since(tr.Start).Milliseconds()) / float64(tr.Duration.Milliseconds())
		if v >= 1 {
			v = 1
			tr.Transitioning = false
			if isCurrent {
				s.removeScene(e, w)
				s.transStart(s.next, -tr.Direction, -tr.InOut)
				s.addScene(s.next, w)
				s.current = nil
			}
			if isNext {
				s.current = e
				s.next = nil
				s.Panel.Refresh(w)
			}
		}
		if isCurrent {
			v = 1 - v
		}
		tr.Apply(e, tr.Ease(v, tr.InOut), tr.Direction)
	}
}

func (s *Scene) UpdateOnce(w golem.World) {
	s.once.Do(func() {
		s.nextScene(w, 0)
	})

	s.Panel.UpdateOnce(w)

	if s.transitioning {
		return
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyRight) ||
		ebiten.IsKeyPressed(ebiten.KeyDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.nextScene(w, 1)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) ||
		ebiten.IsKeyPressed(ebiten.KeyLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.nextScene(w, -1)
	}
}

func (s *Scene) getPanelInfos(w golem.World) string {
	if s.current != nil {
		sc := component.GetScene(s.current)
		if s != nil {
			return fmt.Sprintf("Current scene : %s\n[<]/[>] Prev/Next ", sc.Name)
		}
	}

	return ""
}

func (s *Scene) nextScene(w golem.World, dir int) {
	if len(s.scenes) == 0 {
		return
	}

	// Prevent too fast switch
	if !s.lastSwitch.IsZero() && time.Since(s.lastSwitch) < 200*time.Millisecond {
		return
	}
	s.lastSwitch = time.Now()

	s.idx += dir
	if s.idx < 0 {
		s.idx = len(s.scenes) - 1
	}
	if s.idx >= len(s.scenes) {
		s.idx = 0
	}

	if s.scenes[s.idx] == s.current {
		return
	}

	s.next = s.scenes[s.idx]

	if s.current != nil {
		s.transStart(s.current, float64(dir), component.TransitionOut)
		return
	}

	s.transStart(s.next, float64(-dir), component.TransitionIn)
	s.addScene(s.next, w)
}

func (s *Scene) transStart(e golem.Entity, d, io float64) {
	t := component.GetTransition(e)
	if t != nil && !t.Transitioning {
		t.Start = time.Now()
		t.Transitioning = true
		t.Direction = d
		return
	}

	panic("no transition component, use helpers.TransitionNone to disable transition")
}

func (s *Scene) addScene(e golem.Entity, w golem.World) {
	lf := component.GetLifecycle(e)
	if lf != nil && lf.SetUp != nil {
		lf.SetUp()
	}
	w.AddEntity(e)
}

func (s *Scene) removeScene(e golem.Entity, w golem.World) {
	lf := component.GetLifecycle(e)
	if lf != nil && lf.TearDown != nil {
		lf.TearDown()
	}
	w.RemoveEntity(e)
}
