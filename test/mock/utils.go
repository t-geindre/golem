package mock

import (
	"github.com/t-geindre/golem/pkg/golem"
	"math/rand"
)

const (
	LayerA golem.LayerID = 0
	LayerB
	LayerC
)

func GetSimpleWorld(size int) golem.World {
	w := golem.NewWorld()
	w.AddLayers(LayerA)
	w.AddSystem(NewSysA())
	for i := 0; i < size; i++ {
		w.AddEntity(NewEntityA(LayerA))
	}
	w.Flush()

	return w
}

func GetDeadWorld(size int) golem.World {
	w := golem.NewWorld()
	w.AddLayers(LayerA)
	w.AddSystems(
		NewSysEmpty(), NewSysEmpty(), NewSysEmpty(),
		NewSysEmpty(), NewSysEmpty(), NewSysEmpty(),
	)
	for i := 0; i < size; i++ {
		w.AddEntity(NewEntityA(LayerA))
	}
	w.Flush()

	return w
}

func GetComplexWorld(size int) golem.World {
	layers := []golem.LayerID{LayerA, LayerB, LayerC}
	entities := []func(id golem.LayerID) golem.Entity{
		NewEntityA,
		NewEntityB,
		NewEntityC,
	}

	w := golem.NewWorld()
	w.AddLayers(layers...)
	w.AddSystems(NewSysA(), NewSysB(), NewSysC())

	for i := 0; i < size; i++ {
		l := layers[i%3]
		w.AddEntity(entities[rand.Intn(3)](l))
	}

	w.Flush()

	return w
}
