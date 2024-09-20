package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
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

func NewScene(l golem.LayerID, n string, t component.TransitionFunc, d time.Duration) *Scene {
	s := &Scene{
		Entity:     golem.NewEntity(l),
		World:      golem.NewWorld(),
		Scene:      component.NewScene(n),
		Opacity:    component.NewOpacity(1),
		Scale:      component.NewScale(1, 0.5, 0.5),
		Transition: component.NewTransition(t, d),
		Lifecycle:  component.NewLifecycle(),
	}

	return s
}
