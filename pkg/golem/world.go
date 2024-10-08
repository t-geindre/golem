package golem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"slices"
)

type LayerID uint8

type World interface {
	Clear()
	AddLayers(layers ...LayerID) LayerID
	RemoveLayer(layer LayerID)
	AddEntity(e Entity)
	AddEntities(e ...Entity)
	RemoveEntity(e Entity)
	GetEntities(layer LayerID) []Entity
	SetParentEntity(e Entity)
	GetParentEntity() Entity
	GetLayers() []LayerID
	Size() int
	Flush()
	AddSystem(s System)
	AddSystems(s ...System)
	RemoveSystem(s System)
	Draw(screen *ebiten.Image)
	Update()
	Freeze()
	Unfreeze()
	IsFrozen() bool
}

type world struct {
	layers       []LayerID
	entities     map[LayerID][]Entity
	parent       Entity
	eCount       int
	eChildCount  int
	drawers      []Drawer
	drawersOnce  map[LayerID][]DrawerOnce
	updaters     []Updater
	updatersOnce []UpdaterOnce
	delayed      []func()
	nextLid      LayerID
	frozen       bool
	clk          Clock
}

func NewWorld() World {
	w := &world{}
	w.Clear()
	return w
}

func (w *world) Clear() {
	w.layers = make([]LayerID, 0)

	w.entities = make(map[LayerID][]Entity)

	w.drawers = make([]Drawer, 0)
	w.drawersOnce = make(map[LayerID][]DrawerOnce)

	w.updaters = make([]Updater, 0)
	w.updatersOnce = make([]UpdaterOnce, 0)

	w.delayed = make([]func(), 0)

	w.eCount = 0
	w.eChildCount = 0

	w.clk = newClock()
}

func (w *world) AddLayers(layers ...LayerID) LayerID {
main:
	for _, layer := range layers {
		if layer >= w.nextLid {
			w.nextLid = layer + 1
		}
		for _, l := range w.layers {
			if l == layer {
				continue main
			}
		}
		w.layers = append(w.layers, layer)
		w.entities[layer] = make([]Entity, 0)
		w.drawersOnce[layer] = make([]DrawerOnce, 0)
	}

	return w.nextLid
}

func (w *world) RemoveLayer(layer LayerID) {
	for i, l := range w.layers {
		if l == layer {
			w.layers = append(w.layers[:i], w.layers[i+1:]...)
			delete(w.entities, layer)
			break
		}
	}
}

func (w *world) GetLayers() []LayerID {
	return w.layers
}

func (w *world) AddEntity(e Entity) {
	w.delayed = append(w.delayed, func() {
		w.AddLayers(e.GetLayer())

		e.worldAdded()
		if e.worldCount() > 1 {
			panic("entity already added to another world")
		}

		e.setIndex(len(w.entities[e.GetLayer()]))
		w.entities[e.GetLayer()] = append(w.entities[e.GetLayer()], e)
		w.eCount++
	})
}

func (w *world) AddEntities(e ...Entity) {
	for _, en := range e {
		w.AddEntity(en)
	}
}

func (w *world) RemoveEntity(e Entity) {
	w.delayed = append(w.delayed, func() {
		w.AddLayers(e.GetLayer())
		mln := len(w.entities[e.GetLayer()]) - 1
		idx := e.index()

		if idx < 0 || idx > mln || e != w.entities[e.GetLayer()][idx] {
			// Entity index is not valid, either the entity has already been removed or
			// has been added to another worlds
			return
		}

		e.worldRemoved()
		if idx != mln {
			last := w.entities[e.GetLayer()][mln]
			last.setIndex(idx)
			w.entities[e.GetLayer()][idx] = last
		}
		w.entities[e.GetLayer()] = slices.Delete(w.entities[e.GetLayer()], mln, mln+1)
		w.eCount--
	})
}

func (w *world) GetEntities(layer LayerID) []Entity {
	return w.entities[layer]
}

func (w *world) SetParentEntity(e Entity) {
	w.parent = e
}

func (w *world) GetParentEntity() Entity {
	return w.parent
}

func (w *world) Size() int {
	return w.eCount + w.eChildCount
}

func (w *world) Flush() {
	if len(w.delayed) == 0 {
		return
	}

	for _, do := range w.delayed {
		do()
	}
	w.delayed = w.delayed[:0]
}

func (w *world) AddSystem(s System) {
	if d, ok := s.(Drawer); ok {
		w.drawers = append(w.drawers, d)
	}

	if u, ok := s.(Updater); ok {
		w.updaters = append(w.updaters, u)
	}

	if do, ok := s.(DrawerOnce); ok {
		w.AddLayers(do.GetLayer())
		w.drawersOnce[do.GetLayer()] = append(w.drawersOnce[do.GetLayer()], do)
	}

	if uo, ok := s.(UpdaterOnce); ok {
		w.updatersOnce = append(w.updatersOnce, uo)
	}
}

func (w *world) AddSystems(s ...System) {
	for _, sys := range s {
		w.AddSystem(sys)
	}
}

func (w *world) RemoveSystem(s System) {
	if d, ok := s.(Drawer); ok {
		for i, c := range w.drawers {
			if c == d {
				w.drawers = append(w.drawers[:i], w.drawers[i+1:]...)
				break
			}
		}
	}

	if u, ok := s.(Updater); ok {
		for i, c := range w.updaters {
			if c == u {
				w.updaters = append(w.updaters[:i], w.updaters[i+1:]...)
				break
			}
		}
	}

	if do, ok := s.(DrawerOnce); ok {
		for i, c := range w.drawersOnce[do.GetLayer()] {
			if c == do {
				w.drawersOnce[do.GetLayer()] = append(w.drawersOnce[do.GetLayer()][:i], w.drawersOnce[do.GetLayer()][i+1:]...)
				break
			}
		}
	}

	if uo, ok := s.(UpdaterOnce); ok {
		for i, c := range w.updatersOnce {
			if c == uo {
				w.updatersOnce = append(w.updatersOnce[:i], w.updatersOnce[i+1:]...)
				break
			}
		}
	}
}

func (w *world) Draw(screen *ebiten.Image) {
	for _, layer := range w.layers {
		for _, d := range w.drawersOnce[layer] {
			d.DrawOnce(screen, w)
		}
		for _, e := range w.entities[layer] {
			if sw, ok := e.(World); ok {
				sw.SetParentEntity(e)
				sw.Draw(screen)
				sw.SetParentEntity(nil)
			}
			for _, d := range w.drawers {
				d.Draw(e, screen, w)
			}
		}
	}
}

func (w *world) Update() {
	w.Flush()

	if w.frozen {
		return
	}

	w.clk.Tick()
	eChildCount := 0

	for _, u := range w.updatersOnce {
		u.UpdateOnce(w, w.clk)
	}

	for _, layer := range w.layers {
		for _, e := range w.entities[layer] {
			if sw, ok := e.(World); ok {
				sw.SetParentEntity(e)
				sw.Update()
				sw.SetParentEntity(nil)
				eChildCount += sw.Size()
			}
			for _, u := range w.updaters {
				u.Update(e, w, w.clk)
			}
		}
	}

	w.eChildCount = eChildCount
}

func (w *world) Freeze() {
	w.frozen = true
}

func (w *world) Unfreeze() {
	w.frozen = false
}

func (w *world) IsFrozen() bool {
	return w.frozen
}
