package mock

import (
	"github.com/t-geindre/golem/pkg/golem"
	"math/rand"
)

type EntityC struct {
	golem.Entity
	*CompA
	*CompB
	*CompC
}

func NewEntityC(l golem.LayerID) golem.Entity {
	return &EntityC{
		Entity: golem.NewEntity(l),
		CompA: &CompA{
			Value: rand.Float64(),
		},
		CompB: &CompB{
			Value: rand.Float64(),
		},
		CompC: &CompC{
			Value: rand.Float64(),
		},
	}
}
