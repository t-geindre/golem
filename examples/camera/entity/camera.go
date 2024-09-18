package entity

import (
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type Camera struct {
	golem.Entity
	*component.Position
	*component.Camera
}

func NewCamera(l golem.LayerID, z float64, p image.Point) *Camera {
	return &Camera{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(p),
		Camera:   component.NewCamera(z, z),
	}
}
