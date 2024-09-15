package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/assets"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

func main() {
	const margin = 150
	const wWidth, wHeight = assets.Size + margin*2, assets.Size + margin*2
	const (
		LayerAll = iota
		LayerDebug
	)

	ebiten.SetWindowTitle("Golem example - Scenes")
	ebiten.SetWindowSize(wWidth, wHeight)

	ebiten.SetVsyncEnabled(false)

	g := golemutils.NewGame()
	g.World.AddLayers(LayerAll, LayerDebug)

	scenes := make([]golem.Entity, 0)
	for name, gopher := range map[string]golem.Entity{
		"Normal": entity.NewGopher(LayerAll, margin+assets.Size/2, margin+assets.Size/2),
		"Angry":  entity.NewGopherAngry(LayerAll, margin+assets.Size/2, margin+assets.Size/2),
		"Sad":    entity.NewGopherSad(LayerAll, margin+assets.Size/2, margin+assets.Size/2),
	} {
		scene := entity.NewScene(LayerAll, name)
		scene.AddLayers(LayerAll)
		scene.AddEntity(gopher)
		scene.AddSystem(system.NewRenderer())
		scene.AddSystem(system.NewAnimation())
		scenes = append(scenes, scene)
	}

	g.World.AddSystem(system.NewScene(LayerDebug, scenes...))
	g.World.AddSystem(golemutils.NewMetrics(LayerDebug, time.Millisecond*100))

	ebiten.RunGame(g)
}
