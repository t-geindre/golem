package golem

type EntityID uint64

type Entity interface {
	GetLayer() LayerID
	setIndex(int)
	getIndex() int
}

type entity struct {
	layer LayerID
	idx   int
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

func (e *entity) getIndex() int {
	return e.idx
}
