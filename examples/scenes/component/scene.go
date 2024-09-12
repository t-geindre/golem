package component

import "github.com/t-geindre/golem/pkg/golem"

//go:generate golem component Scene
type Scene struct {
	golem.World
	Name string
}

func NewScene() *Scene {
	return &Scene{
		World: golem.NewWorld(),
	}
}
