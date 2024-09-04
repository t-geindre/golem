package helper

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const size = 20

var Assets []*ebiten.Image

func init() {
	Assets = make([]*ebiten.Image, 0)
	for _, col := range []color.RGBA{
		{0xff, 0x0, 0x0, 0xff},
		{0x0, 0xff, 0x0, 0xff},
		{0x0, 0x0, 0xff, 0xff},
		{0xff, 0xff, 0x00, 0xff},
		{0xff, 0x0, 0xff, 0xff},
		{0x0, 0xff, 0xff, 0xff},
	} {
		img := ebiten.NewImage(size, size)
		vector.StrokeRect(img, 1, 1, size-1, size-1, 1, col, false)
		Assets = append(Assets, img)
	}
}
