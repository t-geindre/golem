package main

import (
	"github.com/t-geindre/golem/pkg/golemutils"
)

type Game struct {
	*golemutils.Game
	ww, wh int
}

func NewGame(ww, wh int) *Game {
	return &Game{
		Game: golemutils.NewGame(),
		ww:   ww,
		wh:   wh,
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ww, g.wh
}
