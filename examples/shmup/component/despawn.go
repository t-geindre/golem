package component

const (
	DespawnDirectionUp = iota
	DespawnDirectionDown
)

//go:generate golem despawn.go
type Despawn struct {
	Direction int
}

func NewDespawn(direction int) *Despawn {
	return &Despawn{
		Direction: direction,
	}
}
