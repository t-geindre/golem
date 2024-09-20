package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/helper"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"golang.org/x/image/colornames"
	"time"
)

const (
	LayerBackground golem.LayerID = iota
	LayerScenes
	LayerDebug
)

func main() {

	ebiten.SetWindowTitle("Golem example - Scenes")
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(false)

	g := golemutils.NewGame()
	g.World.AddLayers(LayerBackground, LayerScenes, LayerDebug)

	g.World.AddSystems(
		system.NewBackground(LayerBackground, colornames.White),
		system.NewFullscreen(),
		system.NewScene(LayerDebug, helper.GetSlides(LayerScenes)...),
		golemutils.NewMetrics(LayerDebug, time.Millisecond*100),
	)

	ebiten.RunGame(g)
}
