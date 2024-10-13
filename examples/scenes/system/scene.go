package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

type SceneBuilder func() golem.Entity

type Scene struct {
	scenes        []golem.Entity
	idx           int
	current, next golem.Entity
	transitioning bool
	lastSwitch    time.Time
	size          int
	*golemutils.Panel
}

func NewScene(l golem.LayerID, scenes ...golem.Entity) *Scene {
	s := &Scene{
		scenes: scenes,
		size:   len(scenes),
	}

	s.Panel = golemutils.NewPanel(l, s.getPanelInfos, golemutils.RefreshOnce)
	s.Panel.Stick = golemutils.StickBottomRight

	return s
}

func (s *Scene) Update(e golem.Entity, w golem.World, _ golem.Clock) {
	tr := component.GetTransition(e)
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
				s.transStart(s.next, -tr.Direction)
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
		tr.Apply(e, tr.Ease(v), tr.Direction)
	}
}

func (s *Scene) UpdateOnce(w golem.World, c golem.Clock) {
	s.Panel.UpdateOnce(w, c)

	if s.current == nil && s.next == nil {
		s.nextScene(w, 0)
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

	s.transitioning = false
}

func (s *Scene) getPanelInfos(w golem.World) string {
	if s.current != nil {
		sc := component.GetScene(s.current)
		if s != nil {
			return fmt.Sprintf("[%d/%d] %s\n[<]/[>] Prev/Next ", s.idx+1, s.size, sc.Name)
		}
	}

	return ""
}

func (s *Scene) nextScene(w golem.World, dir int) {
	if !s.canSwitch() {
		return
	}

	s.idx = ((s.idx + dir%s.size) + s.size) % s.size
	if s.scenes[s.idx] == s.current {
		return
	}
	s.next = s.scenes[s.idx]

	if dir == 0 {
		dir = 1
	}

	if s.current != nil {
		s.transStart(s.current, float64(dir))
		return
	}

	s.transStart(s.next, float64(-dir))
	s.addScene(s.next, w)
}

func (s *Scene) transStart(e golem.Entity, d float64) {
	t := component.GetTransition(e)
	if t != nil && !t.Transitioning {
		t.Start = time.Now()
		t.Transitioning = true
		t.Direction = d
		s.transitioning = true
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

func (s *Scene) canSwitch() bool {
	if len(s.scenes) > 0 &&
		!s.transitioning &&
		(s.lastSwitch.IsZero() || time.Since(s.lastSwitch) > 200*time.Millisecond) {
		s.lastSwitch = time.Now()
		return true
	}
	return false
}
