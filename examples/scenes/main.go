package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/assets"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
)

func main() {
	const margin = 150
	const wWidth, wHeight = assets.Size + margin*2, assets.Size + margin*2
	const LayerAll = iota

	ebiten.SetWindowTitle("Golem example - Scenes")
	ebiten.SetWindowSize(wWidth, wHeight)
	ebiten.SetVsyncEnabled(false)

	w := golem.NewWorld()
	w.AddLayers(LayerAll)

	scenes := make([]golem.Entity, 0)
	for name, gopher := range map[string]golem.Entity{
		"Normal": entity.NewGopher(LayerAll, 0, 0),
		"Angry":  entity.NewGopherAngry(LayerAll, 0, 0),
		"Sad":    entity.NewGopherSad(LayerAll, 0, 0),
	} {
		scene := entity.NewScene(LayerAll, margin+assets.Size/2, margin+assets.Size/2)
		scene.AddLayers(LayerAll)
		scene.AddEntity(gopher)
		scene.AddSystem(system.NewRenderer(scene))
		scene.AddSystem(system.NewAnimation())
		scene.Name = name
		scenes = append(scenes, scene)
	}

	w.AddSystem(system.NewScene(scenes...))

	ebiten.RunGame(NewGame(w))
}
