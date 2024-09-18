package entity

import (
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Camera struct {
	golem.Entity
	*component.Position
	*component.Camera
}

func NewCamera(l golem.LayerID, z, tz float64, px, py int) *Camera {
	return &Camera{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(px, py),
		Camera:   component.NewCamera(z, tz),
	}
}
