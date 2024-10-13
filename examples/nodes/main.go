package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/nodes/entity"
	"github.com/t-geindre/golem/examples/nodes/system"
	"github.com/t-geindre/golem/pkg/golemutils"
	"math/rand"
	"time"
)

func main() {
	const (
		LayerAll = iota
		LayerDebug
	)

	ww, wh := 800, 800
	ebiten.SetWindowSize(ww, wh)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	root := entity.NewNode(LayerAll, 1, 2, 30, 30, ww-60, wh-60)
	for i := 0; i < 10; i++ {
		cont := entity.NewNode(LayerAll,
			1+rand.Intn(5), 1+rand.Intn(5),
			rand.Intn(ww-60), rand.Intn(wh-60),
			200, 200)
		for j := 0; j < 10; j++ {
			child := entity.NewNode(LayerAll,
				rand.Intn(3), rand.Intn(3),
				rand.Intn(190), rand.Intn(190),
				10, 10)
			cont.AddEntity(child)
		}
		root.AddEntity(cont)
	}

	game := golemutils.NewGame()
	game.AddLayers(LayerAll, LayerDebug)
	game.AddEntity(root)

	game.AddSystems(
		system.NewMove(),
		system.NewBounce(),
		system.NewGeometry(),
		system.NewRenderer(),
		golemutils.NewMetrics(LayerDebug, time.Millisecond*100),
	)

	ebiten.RunGame(game)
}
