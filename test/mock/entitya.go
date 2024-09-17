package mock

import (
	"github.com/t-geindre/golem/pkg/golem"
	"math/rand"
)

type EntityA struct {
	golem.Entity
	*CompA
}

func NewEntityA(l golem.LayerID) golem.Entity {
	return &EntityA{
		Entity: golem.NewEntity(l),
		CompA: &CompA{
			Value: rand.Float64(),
		},
	}
}
