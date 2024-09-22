package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Scene struct {
	golem.Entity
	golem.World
	*component.Scene
	*component.Opacity
	*component.Scale
	*component.Transition
	*component.Lifecycle
}

func NewScene(l golem.LayerID, name string) *Scene {
	s := &Scene{
		Entity:  golem.NewEntity(l),
		World:   golem.NewWorld(),
		Scene:   component.NewScene(name),
		Opacity: component.NewOpacity(1),
		Scale:   component.NewScale(1, .5, .5),
	}

	return s
}
