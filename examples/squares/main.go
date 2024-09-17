package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/squares/system"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

func main() {
	const (
		LayerSquares = iota
		LayerDebug
	)

	ebiten.SetWindowTitle("Golem example - Squares")
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(false)

	g := golemutils.NewGame()

	g.World.AddLayers(LayerSquares, LayerDebug)

	g.World.AddSystems(
		system.NewBounce(),
		system.NewMove(),
		system.NewRenderer(),
		golemutils.NewMetrics(LayerDebug, time.Millisecond*100),
		system.NewSpawner(LayerSquares, LayerDebug),
	)

	ebiten.RunGame(g)
}
