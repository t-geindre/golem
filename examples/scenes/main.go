package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/helper"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"image/color"
	"time"
)

const (
	LayerBackground golem.LayerID = iota
	LayerScenes
	LayerDebug
)

func main() {
	mw, mh := ebiten.Monitor().Size()

	ebiten.SetWindowTitle("Golem example - Scenes")
	ebiten.SetWindowSize(mw, mh)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(false)

	g := NewGame(mw, mh)
	g.World.AddLayers(LayerBackground, LayerScenes, LayerDebug)

	g.World.AddSystems(
		system.NewBackground(LayerBackground, color.RGBA{R: 0xf4, G: 0xf9, B: 0xff, A: 0xff}),
		system.NewFullscreen(),
		system.NewScene(LayerDebug, helper.GetSlides(LayerScenes)...),
		golemutils.NewMetrics(LayerDebug, time.Millisecond*100),
	)

	ebiten.RunGame(g)
}
