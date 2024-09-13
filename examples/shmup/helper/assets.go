package helper

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/t-geindre/golem/examples/shmup/assets"
	"image"
)

var Assets map[string]*ebiten.Image

const Scale = 2
const TileSize = 16 * Scale

func init() {

	Assets = make(map[string]*ebiten.Image)

	plSheet := getAssetImage("player_16.png")
	Assets["player"] = getTile(plSheet, 1, 0)

	enSheet := getAssetImage("enemy_16.png")
	Assets["enemy_f1"] = getTile(enSheet, 0, 0)
	Assets["enemy_f2"] = getTile(enSheet, 1, 0)
	Assets["enemy_f3"] = getTile(enSheet, 2, 0)
	Assets["enemy_f4"] = getTile(enSheet, 3, 0)
	Assets["enemy_f5"] = getTile(enSheet, 4, 0)

	bulletSheet := getAssetImage("bullet_16.png")
	Assets["bullet_f1"] = getTile(bulletSheet, 0, 0)
	Assets["bullet_f2"] = getTile(bulletSheet, 1, 0)
}

func getAssetImage(file string) *ebiten.Image {
	b, e := assets.Assets.ReadFile(file)
	if e != nil {
		panic(e)
	}

	i, _, e := ebitenutil.NewImageFromReader(bytes.NewReader(b))
	if e != nil {
		panic(e)
	}

	return scaleImage(i, Scale)
}

func scaleImage(img *ebiten.Image, scale float64) *ebiten.Image {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	sw, sh := int(float64(w)*scale), int(float64(h)*scale)
	scaled := ebiten.NewImage(sw, sh)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	scaled.DrawImage(img, op)
	return scaled
}

func getTile(img *ebiten.Image, x, y int) *ebiten.Image {
	return img.SubImage(image.Rect(x*TileSize, y*TileSize, (x+1)*TileSize, (y+1)*TileSize)).(*ebiten.Image)
}
