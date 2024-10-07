package golem

type EntityID uint64

type Entity interface {
	GetLayer() LayerID
	setIndex(int)
	index() int
	worldAdded()
	worldRemoved()
	worldCount() int8
}

type entity struct {
	layer  LayerID
	idx    int
	worlds int8
}

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

func (e *entity) worldAdded() {
	e.worlds++
}

func (e *entity) worldRemoved() {
	e.worlds--
}

func (e *entity) worldCount() int8 {
	return e.worlds
}
