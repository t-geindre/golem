package component

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
