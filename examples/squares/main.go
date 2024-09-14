package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/squares/system"
	"github.com/t-geindre/golem/pkg/golemutils"
	"image"
	"time"
)

func main() {
	winSize := &image.Point{X: 1000, Y: 1000}

	const (
		LayerSquares = iota
		LayerDebug
	)

	ebiten.SetWindowTitle("Golem example - Squares")
	ebiten.SetWindowSize(winSize.X, winSize.Y)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(false)

	g := NewGame(winSize)

	g.World.AddLayers(LayerSquares, LayerDebug)

	g.World.AddSystem(system.NewBounce(winSize))
	g.World.AddSystem(system.NewMove())
	g.World.AddSystem(system.NewRenderer())
	g.World.AddSystem(golemutils.NewMetrics(LayerDebug, time.Millisecond*100))
	g.World.AddSystem(system.NewSpawner(LayerSquares, LayerDebug, winSize))

	ebiten.RunGame(g)
}
