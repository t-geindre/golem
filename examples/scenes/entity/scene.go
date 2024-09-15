package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Scene struct {
	golem.Entity
	golem.World
	*component.Scene
	*component.Opacity
	*component.Transition
}

func NewScene(layer golem.LayerID, name string) *Scene {
	s := &Scene{
		Entity:     golem.NewEntity(layer),
		World:      golem.NewWorld(),
		Scene:      component.NewScene(name),
		Opacity:    component.NewOpacity(1),
		Transition: component.NewTransition(helper.TransitionFade, time.Millisecond*250),
	}

	return s
}
