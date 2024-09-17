package mock

import (
	"github.com/t-geindre/golem/pkg/golem"
	"math/rand"
)

type EntityB struct {
	golem.Entity
	*CompA
	*CompB
}

func NewEntityB(l golem.LayerID) golem.Entity {
	return &EntityB{
		Entity: golem.NewEntity(l),
		CompA: &CompA{
			Value: rand.Float64(),
		},
		CompB: &CompB{
			Value: rand.Float64(),
		},
	}
}
