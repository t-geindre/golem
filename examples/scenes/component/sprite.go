package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

//go:generate golem component Sprite
type Sprite struct {
	Img *ebiten.Image
}

func NewSprite(img *ebiten.Image) *Sprite {
	return &Sprite{
		Img: img,
	}
}
