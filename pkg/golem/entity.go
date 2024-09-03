package golem

type EntityID uint64

type Entity interface {
	GetID() EntityID
	setId(id EntityID)
	GetLayer() LayerID
}

type entity struct {
	id    EntityID
	layer LayerID
}

func NewEntity(layer LayerID) Entity {
	return &entity{layer: layer}
}

func (e *entity) GetID() EntityID {
	return e.id
}

func (e *entity) setId(id EntityID) {
	e.id = id
}

func (e *entity) GetLayer() LayerID {
	return e.layer
}
