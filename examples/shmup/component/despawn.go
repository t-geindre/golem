package component

const (
	DespawnDirectionUp = iota
	DespawnDirectionDown
)

//go:generate golem component Despawn
type Despawn struct {
	Direction int
}

func NewDespawn(direction int) *Despawn {
	return &Despawn{
		Direction: direction,
	}
}
