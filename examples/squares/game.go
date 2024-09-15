package main

import (
	"github.com/t-geindre/golem/pkg/golemutils"
	"image"
)

type Game struct {
	*golemutils.Game
	rect *image.Point
}

func NewGame(rect *image.Point) *Game {
	return &Game{Game: golemutils.NewGame(), rect: rect}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// We need our own ebiten.Game implementation to keep track of the window size
	// TODO use ebiten.WindowSize() instead
	g.rect.X, g.rect.Y = outsideWidth, outsideHeight
	return outsideWidth, outsideHeight
}
