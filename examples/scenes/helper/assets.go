package helper

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/t-geindre/golem/examples/scenes/assets"
	"image"
)

var Assets []*ebiten.Image

func init() {
	Assets = make([]*ebiten.Image, 0)
	src, _, e := ebitenutil.NewImageFromReader(bytes.NewReader(assets.Gophers))
	if e != nil {
		panic(e)
	}

	for x := 0; x < 7; x++ {
		for y := 0; y < 5; y++ {
			Assets = append(Assets, src.SubImage(image.Rectangle{
				Min: image.Point{X: x * assets.Size, Y: y * assets.Size},
				Max: image.Point{X: (x + 1) * assets.Size, Y: (y + 1) * assets.Size},
			}).(*ebiten.Image))
		}
	}
}
