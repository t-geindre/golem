package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/camera/assets"
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/examples/camera/entity"
	"github.com/t-geindre/golem/examples/camera/helper"
	"github.com/t-geindre/golem/examples/camera/system"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

func main() {
	ebiten.SetWindowTitle("Golem example - Camera")
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(false)

	currentLayer := golem.LayerID(0)
	g := golemutils.NewGame()

	tileset := helper.NewTileset(helper.LoadImage(assets.Tileset), 27, 16, 16)
	m := helper.LoadMap(assets.Map)
	for _, layer := range m.Layers {
		currentLayer++
		g.World.AddLayers(currentLayer)

		for i, tile := range layer.Data {
			if tile == 0 {
				continue
			}
			e := entity.NewTile(
				currentLayer, i%m.Width*16, i/m.Height*16,
				component.NewFrame(tileset.GetTile(tile), time.Second/2),
			)
			g.World.AddEntity(e)
		}
	}

	currentLayer++
	g.World.AddLayers(currentLayer)

	g.World.AddSystems(
		system.NewAnimation(),
		system.NewCamera(),
		golemutils.NewMetrics(currentLayer, time.Millisecond*100),
	)

	ebiten.RunGame(g)
}
