package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/helper"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	LayerBackground golem.LayerID = iota
	LayerScenes
	LayerDebug
)

func main() {
	filePath := "embd://slides.xml"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	if helper.FileExists("slides.xml") {
		filePath = "slides.xml"
	}

	file, err := helper.OpenFile(filePath)
	if err != nil {
		panic(err)
	}

	err = os.Chdir(filepath.Dir(
		strings.Replace(filePath, "embd://", "", 1),
	))
	if err != nil {
		panic(err)
	}

	xml, err := helper.ParseXML(file)
	_ = file.Close()
	if err != nil {
		panic(err)
	}

	loader := helper.NewSlideLoader(LayerScenes)
	err = loader.LoadXML(xml)
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowTitle("Golem example - Scenes")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	slW, slH := loader.GetLayout()
	ebiten.SetWindowSize(int(slW), int(slH))

	g := golemutils.NewGame()
	g.AddLayers(LayerBackground, LayerScenes, LayerDebug)

	g.AddSystems(
		system.NewBackground(LayerBackground, loader.GetBackgroundColor()),
		system.NewFullscreen(),
		system.NewMouse(),
		system.NewScene(LayerDebug, loader.GetSlides(LayerScenes)...),
		golemutils.NewMetrics(LayerDebug, time.Millisecond*100),
	)

	_ = ebiten.RunGame(g)
}
