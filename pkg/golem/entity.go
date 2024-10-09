package golem

type entity struct {
	layer LayerID
	idx   int
	world bool
}

// NewEntity creates a new Golem entity with the given layer
// Any entity added to a world must be an Entity
func NewEntity(l LayerID) Entity {
	return &entity{layer: l}
}

func (e *entity) GetLayer() LayerID {
	return e.layer
}

func (e *entity) setIndex(idx int) {
	e.idx = idx
}

func (e *entity) index() int {
	return e.idx
}

func (e *entity) hasWorld() bool {
	ret := e.world
	e.world = !e.world

	return ret
}
