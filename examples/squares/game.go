package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type Game struct {
	w    golem.World
	rect *image.Point
}

func NewGame(w golem.World, rect *image.Point) *Game {
	return &Game{w: w, rect: rect}
}

func (g *Game) Update() error {
	g.w.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.w.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.rect.X, g.rect.Y = outsideWidth, outsideHeight
	return outsideWidth, outsideHeight
}
