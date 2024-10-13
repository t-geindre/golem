package system

import (
	"github.com/t-geindre/golem/examples/nodes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Geometry struct {
}

func NewGeometry() *Geometry {
	return &Geometry{}
}

func (g *Geometry) Update(e golem.Entity, w golem.World, _ golem.Clock) {
	geo := component.GetGeometry(e)
	if geo == nil {
		return
	}

	geo.GeoM.Reset()

	parent := w.GetParentEntity()
	if parent != nil {
		if parentGeo := component.GetGeometry(parent); parentGeo != nil {
			geo.GeoM.Concat(*parentGeo.GeoM)
		}
	}

	bds := component.GetBoundaries(e)
	if bds == nil {
		return
	}

	geo.Translate(float64(bds.Rectangle.Min.X), float64(bds.Rectangle.Min.Y))
}
