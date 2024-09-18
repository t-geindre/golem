package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/camera/entity"
	"github.com/t-geindre/golem/examples/camera/helper"
	"github.com/t-geindre/golem/examples/camera/system"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

func main() {
	ebiten.SetWindowTitle("Golem example - Camera")
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(false)

	g := golemutils.NewGame()
	m := helper.NewMapFromFile("map3.tmj")

	nextLayer := g.World.AddLayers(m.Layers()...)
	g.World.AddEntities(m.Entities()...)

	cam := entity.NewCamera(nextLayer, 2, m.Center())
	g.World.AddEntity(cam)

	g.World.AddSystems(
		system.NewAnimation(),
		system.NewCamera(),
		system.NewRenderer(cam),
		golemutils.NewMetrics(nextLayer, time.Millisecond*100),
	)

	ebiten.RunGame(g)
}
