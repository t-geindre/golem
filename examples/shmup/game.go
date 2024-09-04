package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
)

type Game struct {
	w golem.World
}

func NewGame(w golem.World) *Game {
	return &Game{w: w}
}

func (g *Game) Update() error {
	g.w.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.w.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
