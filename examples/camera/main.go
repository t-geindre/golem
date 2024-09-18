package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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

	tileset := helper.NewTilesetFromFile("tileset.tsx")
	tw, th := tileset.GetTileSize()
	m := helper.LoadMapFromFile("map3.tmj")

	for _, layer := range m.Layers {
		currentLayer++
		g.World.AddLayers(currentLayer)

		for i, tile := range layer.Data {
			if tile == 0 {
				continue
			}
			t := tileset.GetTile(tile)
			fs := make([]component.Frame, 0)
			if t.Animation != nil {
				for _, f := range t.Animation {
					fs = append(fs, component.Frame{Img: f.Img, Duration: f.Duration})
				}
			} else {
				fs = append(fs, component.Frame{Img: t.Img})
			}
			e := entity.NewTile(currentLayer, i%m.Width*tw, i/m.Height*th, fs...)
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
