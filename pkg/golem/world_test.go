package golem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestWorldAddRemoveEntity(t *testing.T) {
	w := NewWorld()
	layers := []LayerID{0, 1, 2, 3}
	enCount := 10
	ens := make([]Entity, 0, len(layers)*enCount)

	for _, l := range layers {
		for i := 0; i < enCount; i++ {
			e := NewEntity(l)
			w.AddEntity(e)

			ens = append(ens, e)
		}
		if len(w.GetEntities(l)) != 0 {
			t.Errorf("world not flushed, entity should not be in any layer yet")
		}
	}

	w.Flush()

	for _, l := range layers {
		if len(w.GetEntities(l)) != enCount {
			t.Errorf("entity should be in layer")
		}
	}

	expectedSize := len(layers) * enCount
	if w.Size() != expectedSize {
		t.Errorf("world size is not correct, %d expected, got %d", expectedSize, w.Size())
	}

	for _, e := range ens {
		w.RemoveEntity(e)
	}

	w.Flush()

	if w.Size() != 0 {
		t.Errorf("world should be empty")
	}

	for _, l := range layers {
		if len(w.GetEntities(l)) != 0 {
			t.Errorf("layers should be empty")
		}
	}
}

func TestWorldMultipleRemoval(t *testing.T) {
	w := NewWorld()
	e := NewEntity(0)
	w.AddEntity(e)
	w.Flush()

	w.RemoveEntity(e)
	w.RemoveEntity(e)
	w.Flush()

	if w.Size() != 0 {
		t.Errorf("world should be empty")
	}
}

func TestWorldSystemsCalls(t *testing.T) {
	w := NewWorld()
	enCount, sysCount := 10, 5
	systems := make([]*SystemTracker, 0, sysCount)

	for i := 0; i < sysCount; i++ {
		sys := &SystemTracker{}
		w.AddSystem(sys)
		systems = append(systems, sys)
	}

	for i := 0; i < enCount; i++ {
		e := NewEntity(0)
		w.AddEntity(e)
	}
	w.Flush()

	w.Update()
	for _, sys := range systems {
		if sys.UOC != 1 {
			t.Errorf("UpdateOnce count is not correct, 1 expected, got %d", sys.UOC)
		}
		if sys.UC != enCount {
			t.Errorf("Update count is not correct, %d expected, got %d", enCount, sys.UC)
		}
	}

	w.Draw(&ebiten.Image{})
	for _, sys := range systems {
		if sys.DOC != 1 {
			t.Errorf("DrawOnce count is not correct, 1 expected, got %d", sys.DOC)
		}
		if sys.DC != enCount {
			t.Errorf("Draw count is not correct, %d expected, got %d", enCount, sys.DC)
		}
	}
}

func TestWorldSystemsRemoval(t *testing.T) {
	w := NewWorld()
	sys := &SystemTracker{}
	w.AddSystem(sys)
	w.Update()

	if sys.UOC != 1 {
		t.Errorf("UpdateOnce count is not correct, 1 expected, got %d", sys.UOC)
	}

	w.RemoveSystem(sys)
	w.Update()
	if sys.UOC != 1 {
		t.Errorf("UpdateOnce count is not correct, 1 expected, got %d", sys.UOC)
	}

}

func TestWordEmbeddedWorldDrawUpdate(t *testing.T) {
	en := &struct {
		Entity
		World
	}{
		Entity: NewEntity(0),
		World:  NewWorld(),
	}

	subSys := &SystemTracker{}
	en.AddSystem(subSys)
	en.AddEntity(NewEntity(0))
	en.Flush()

	w := NewWorld()
	sys := &SystemTracker{}
	w.AddSystems(sys)
	w.AddEntity(en)
	w.Flush()

	w.Update()
	w.Draw(&ebiten.Image{})

	if sys.DC != 1 || sys.DOC != 1 || sys.UC != 1 || sys.UOC != 1 {
		t.Errorf("root world should be updated")
	}

	if subSys.DC != 1 || subSys.DOC != 1 || subSys.UC != 1 || subSys.UOC != 1 {
		t.Errorf("embedded world should be updated")
	}

	if subSys.LastParent != en {
		t.Errorf("embedded world should have the parent entity")
	}
}

func TestWorldClear(t *testing.T) {
	w := NewWorld()
	enCount := 10
	for i := 0; i < enCount; i++ {
		w.AddEntity(NewEntity(0))
	}
	w.Flush()

	w.Clear()

	if w.Size() != 0 {
		t.Errorf("world should be empty")
	}
}

func TestWorldClearAndAddAgain(t *testing.T) {
	w := NewWorld()
	e := NewEntity(0)
	w.AddEntity(e)
	w.Flush()

	w.Clear()
	w.AddEntity(e)
	w.Flush()

	if w.Size() != 1 {
		t.Errorf("world should have one entity")
	}
}

func TestWorldPanicOnAddEntityAlreadyInAnotherWorld(t *testing.T) {
	w1, w2 := NewWorld(), NewWorld()
	e := NewEntity(0)
	w1.AddEntity(e)
	w1.Flush()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should panic")
		}
	}()

	w2.AddEntity(e)
	w2.Flush()
}

type SystemTracker struct {
	UC, UOC, DC, DOC, GLC int
	LastParent            Entity
}

func (s *SystemTracker) Update(Entity, World, Clock) {
	s.UC++
}

func (s *SystemTracker) UpdateOnce(w World, _ Clock) {
	if w.GetParentEntity() != nil {
		s.LastParent = w.GetParentEntity()
	}

	s.UOC++
}

func (s *SystemTracker) Draw(Entity, *ebiten.Image, World) {
	s.DC++
}

func (s *SystemTracker) DrawOnce(*ebiten.Image, World) {
	s.DOC++
}

func (s *SystemTracker) GetLayer() LayerID {
	s.GLC++
	return 0
}
