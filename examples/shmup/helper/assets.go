package helper

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

var Assets map[string]*ebiten.Image

func init() {
	Assets = make(map[string]*ebiten.Image)
	Assets["player"] = buildTriangleImg(21, 21, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	Assets["enemy"] = buildSquareImg(20, 20, color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff})
	Assets["bullet"] = buildSquareImg(6, 6, color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff})
}

func buildTriangleImg(w, h float32, col color.RGBA) *ebiten.Image {
	img := ebiten.NewImage(int(w), int(h))
	vector.StrokeLine(img, 1, h-1, w-1, h-1, 1, col, true)
	vector.StrokeLine(img, 1, h-1, w/2, 1, 1, col, true)
	vector.StrokeLine(img, w-1, h-1, w/2, 1, 1, col, true)
	return img
}

func buildSquareImg(w, h float32, col color.RGBA) *ebiten.Image {
	img := ebiten.NewImage(int(w), int(h))
	vector.StrokeRect(img, 1, 1, w-1, h-1, 1, col, false)
	return img
}
