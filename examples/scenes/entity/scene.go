package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	*component.Boundaries
	*component.Position
}

func NewScene(l golem.LayerID, name string) *Scene {
	mw, mh := ebiten.Monitor().Size()
	s := &Scene{
		Entity:     golem.NewEntity(l),
		World:      golem.NewWorld(),
		Scene:      component.NewScene(name),
		Opacity:    component.NewOpacity(1),
		Scale:      component.NewScale(1, .5, .5),
		Boundaries: component.NewBoundaries(0, 0, mw, mh),
		Position:   component.NewPosition(.5, .5, .5, .5),
	}

	return s
}
