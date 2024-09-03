package golem

import "github.com/hajimehoshi/ebiten/v2"

type LayerID uint8

type World interface {
	Clear()
	AddLayers(layers ...LayerID)
	RemoveLayer(layer LayerID)
	AddEntity(e Entity)
	RemoveEntity(e Entity)
	Flush()
	AddSystem(s System)
	RemoveSystem(s System)
	Draw(screen *ebiten.Image)
	Update()
}

type world struct {
	layers       []LayerID
	ids          EntityID
	entities     map[LayerID][]Entity
	drawers      []Drawer
	drawersOnce  map[LayerID][]DrawerOnce
	updaters     []Updater
	updatersOnce []UpdaterOnce
	delayed      []func()
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
}

func (w *world) AddLayers(layers ...LayerID) {
main:
	for _, layer := range layers {
		for _, l := range w.layers {
			if l == layer {
				continue main
			}
		}
		w.layers = append(w.layers, layer)
		w.entities[layer] = make([]Entity, 0)
		w.drawersOnce[layer] = make([]DrawerOnce, 0)
	}
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

func (w *world) AddEntity(e Entity) {
	w.delayed = append(w.delayed, func() {
		w.ids++
		e.setId(w.ids)
		w.AddLayers(e.GetLayer())
		w.entities[e.GetLayer()] = append(w.entities[e.GetLayer()], e)
	})
}

func (w *world) RemoveEntity(e Entity) {
	w.delayed = append(w.delayed, func() {
		for i, candidate := range w.entities[e.GetLayer()] {
			if candidate.GetID() == e.GetID() {
				w.entities[e.GetLayer()] = append(w.entities[e.GetLayer()][:i], w.entities[e.GetLayer()][i+1:]...)
				break
			}
		}
	})
}

func (w *world) Flush() {
	for _, do := range w.delayed {
		do()
	}
	w.delayed = w.delayed[:0]
}

func (w *world) AddSystem(s System) {
	switch sys := s.(type) {
	case Drawer:
		w.drawers = append(w.drawers, sys)
	case Updater:
		w.updaters = append(w.updaters, sys)
	case DrawerOnce:
		w.AddLayers(sys.GetLayer())
		w.drawersOnce[sys.GetLayer()] = append(w.drawersOnce[sys.GetLayer()], sys)
	case UpdaterOnce:
		w.updatersOnce = append(w.updatersOnce, sys)
	default:
		panic("system must implement Drawer or Updater")
	}
}

func (w *world) RemoveSystem(s System) {
	switch sys := s.(type) {
	case Drawer:
		for i, d := range w.drawers {
			if d == sys {
				w.drawers = append(w.drawers[:i], w.drawers[i+1:]...)
				break
			}
		}
	case Updater:
		for i, u := range w.updaters {
			if u == sys {
				w.updaters = append(w.updaters[:i], w.updaters[i+1:]...)
				break
			}
		}
	case DrawerOnce:
		for i, d := range w.drawersOnce[sys.GetLayer()] {
			if d == sys {
				w.drawersOnce[sys.GetLayer()] = append(w.drawersOnce[sys.GetLayer()][:i], w.drawersOnce[sys.GetLayer()][i+1:]...)
				break
			}
		}
	case UpdaterOnce:
		for i, u := range w.updatersOnce {
			if u == sys {
				w.updatersOnce = append(w.updatersOnce[:i], w.updatersOnce[i+1:]...)
				break
			}

		}
	default:
		panic("system must implement Drawer or Updater")
	}
}

func (w *world) Draw(screen *ebiten.Image) {
	for _, layer := range w.layers {
		for _, d := range w.drawersOnce[layer] {
			d.Draw(screen, w)
		}
		for _, e := range w.entities[layer] {
			for _, d := range w.drawers {
				d.Draw(e, screen, w)
			}
		}
	}
}

func (w *world) Update() {
	w.Flush()

	for _, layer := range w.layers {
		for _, e := range w.entities[layer] {
			for _, u := range w.updaters {
				u.Update(e, w)
			}
		}
	}
}
