package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
	"image/color"
)

type Background struct {
	color color.RGBA
	layer golem.LayerID
}

func NewBackground(l golem.LayerID, c color.RGBA) *Background {
	return &Background{
		color: c,
		layer: l,
	}
}

func (b *Background) DrawOnce(screen *ebiten.Image, w golem.World) {
	screen.Fill(b.color)
}

func (b *Background) GetLayer() golem.LayerID {
	return b.layer
}
