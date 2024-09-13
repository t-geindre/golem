package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Scene struct {
	golem.Entity
	*component.Scene
	*component.Position
}

func NewScene(layer golem.LayerID, x, y float64) *Scene {
	return &Scene{
		Entity:   golem.NewEntity(layer),
		Scene:    component.NewScene(),
		Position: component.NewPosition(x, y),
	}
}
