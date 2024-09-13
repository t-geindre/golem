package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/squares/system"
	"github.com/t-geindre/golem/pkg/golem"
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

	w := golem.NewWorld()

	w.AddLayers(LayerSquares, LayerDebug)

	w.AddSystem(system.NewBounce(winSize))
	w.AddSystem(system.NewMove())
	w.AddSystem(system.NewRenderer())
	w.AddSystem(golemutils.NewMetrics(LayerDebug, time.Millisecond*100))
	w.AddSystem(system.NewSpawner(LayerSquares, LayerDebug, winSize))

	ebiten.RunGame(NewGame(w, winSize))
}
