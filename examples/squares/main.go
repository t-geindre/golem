package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/squares/entity"
	"github.com/t-geindre/golem/examples/squares/helper"
	"github.com/t-geindre/golem/examples/squares/system"
	"github.com/t-geindre/golem/pkg/golem"
	"math/rand"
)

func main() {
	const wWidth, wHeight = 1000, 1000
	const nbSquares = 5000

	const (
		LayerSquares = iota
		LayerDebug
	)

	ebiten.SetWindowTitle("Golem example - Squares")
	ebiten.SetWindowSize(wWidth, wHeight)
	ebiten.SetVsyncEnabled(false)

	w := golem.NewWorld()

	w.AddLayers(LayerSquares, LayerDebug)

	w.AddSystem(system.NewBounce(wWidth, wHeight))
	w.AddSystem(system.NewMove())
	w.AddSystem(system.NewRenderer())
	w.AddSystem(system.NewDebug(LayerDebug))

	for i := 0; i < nbSquares; i++ {
		s := entity.NewSquare(
			LayerSquares,
			helper.Assets[i%len(helper.Assets)],
			rand.Float64()*wWidth, rand.Float64()*wHeight,
			rand.Float64()*2-1, rand.Float64()*2-1,
		)

		w.AddEntity(s)
	}

	ebiten.RunGame(NewGame(w))
}
