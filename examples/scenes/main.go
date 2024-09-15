package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/examples/scenes/helper"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

func main() {
	const (
		LayerAll = iota
		LayerDebug
	)

	ebiten.SetWindowTitle("Golem example - Scenes")
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(false)

	g := golemutils.NewGame()
	g.World.AddLayers(LayerAll, LayerDebug)

	scenes := make([]golem.Entity, 0)
	for _, scene := range []struct {
		name     string
		frames   []component.Frame
		trans    func(entity golem.Entity, v float64)
		duration time.Duration
	}{
		{
			"Fade transition",
			[]component.Frame{
				component.NewFrame(helper.Assets[0], time.Second*3),
				component.NewFrame(helper.Assets[1], time.Millisecond*200),
			},
			helper.TransitionFade,
			time.Millisecond * 250,
		},
		{
			"Scale transition",
			[]component.Frame{
				component.NewFrame(helper.Assets[15], time.Second*3),
				component.NewFrame(helper.Assets[20], time.Second*3),
				component.NewFrame(helper.Assets[1], time.Millisecond*200),
				component.NewFrame(helper.Assets[20], time.Second*3),
				component.NewFrame(helper.Assets[1], time.Millisecond*200),
			},
			helper.TransitionScale,
			time.Millisecond * 200,
		},
		{
			"No transition",
			[]component.Frame{
				component.NewFrame(helper.Assets[22], time.Second*3),
				component.NewFrame(helper.Assets[13], time.Second*1),
				component.NewFrame(helper.Assets[18], time.Second*1),
				component.NewFrame(helper.Assets[13], time.Second*1),
			},
			helper.TransitionNone,
			0,
		},
	} {
		se := entity.NewScene(LayerAll, scene.name, scene.trans, scene.duration)
		se.Lifecycle.SetUp = func() {
			se.World.AddLayers(LayerAll)
			se.World.AddEntity(entity.NewGopher(LayerAll, scene.frames...))
			se.AddSystem(system.NewRenderer())
			se.AddSystem(system.NewAnimation())
		}
		se.Lifecycle.TearDown = func() {
			se.World.Clear()
		}
		scenes = append(scenes, se)
	}

	g.World.AddSystem(system.NewScene(LayerDebug, scenes...))
	g.World.AddSystem(golemutils.NewMetrics(LayerDebug, time.Millisecond*100))

	ebiten.RunGame(g)
}
