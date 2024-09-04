package component

import "github.com/t-geindre/golem/pkg/golem"

type Despawn interface {
	GetDespawn() *DespawnImpl
}

const (
	DespawnDirectionUp = iota
	DespawnDirectionDown
)

type DespawnImpl struct {
	Direction int
}

func NewDespawn(direction int) *DespawnImpl {
	return &DespawnImpl{
		Direction: direction,
	}
}

func (d *DespawnImpl) GetDespawn() *DespawnImpl {
	return d
}

func GetDespawn(e golem.Entity) *DespawnImpl {
	if d, ok := e.(Despawn); ok {
		return d.GetDespawn()
	}
	return nil
}
