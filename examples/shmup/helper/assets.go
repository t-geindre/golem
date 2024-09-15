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
	sheet, _, e := ebitenutil.NewImageFromReader(bytes.NewReader(assets.Sheet))
	if e != nil {
		panic(e)
	}
	sheet = scaleImage(sheet, Scale)

	Assets = make(map[string]*ebiten.Image)

	Assets["player_f1"] = getTile(sheet, 2, 0, TileSize, TileSize*2)
	Assets["player_f2"] = getTile(sheet, 3, 0, TileSize, TileSize*2)

	Assets["en_lips_f1"] = getTile(sheet, 0, 4, TileSize, TileSize)
	Assets["en_lips_f2"] = getTile(sheet, 1, 4, TileSize, TileSize)
	Assets["en_lips_f3"] = getTile(sheet, 2, 4, TileSize, TileSize)
	Assets["en_lips_f4"] = getTile(sheet, 3, 4, TileSize, TileSize)
	Assets["en_lips_f5"] = getTile(sheet, 4, 4, TileSize, TileSize)

	Assets["en_allan_f1"] = getTile(sheet, 0, 6, TileSize, TileSize)
	Assets["en_allan_f2"] = getTile(sheet, 1, 6, TileSize, TileSize)
	Assets["en_allan_f3"] = getTile(sheet, 2, 6, TileSize, TileSize)
	Assets["en_allan_f4"] = getTile(sheet, 3, 6, TileSize, TileSize)
	Assets["en_allan_f5"] = getTile(sheet, 4, 6, TileSize, TileSize)
	Assets["en_allan_f6"] = getTile(sheet, 4, 6, TileSize, TileSize)

	Assets["en_bonbon_f1"] = getTile(sheet, 0, 5, TileSize, TileSize)
	Assets["en_bonbon_f2"] = getTile(sheet, 1, 5, TileSize, TileSize)
	Assets["en_bonbon_f3"] = getTile(sheet, 2, 5, TileSize, TileSize)
	Assets["en_bonbon_f4"] = getTile(sheet, 3, 5, TileSize, TileSize)

	Assets["bullet_f1"] = getTile(sheet, 4, 5, TileSize, TileSize)
	Assets["bullet_f2"] = getTile(sheet, 5, 5, TileSize, TileSize)

	Assets["explosion_f1"] = getTile(sheet, 0, 3, TileSize, TileSize)
	Assets["explosion_f2"] = getTile(sheet, 1, 3, TileSize, TileSize)
	Assets["explosion_f3"] = getTile(sheet, 2, 3, TileSize, TileSize)
	Assets["explosion_f4"] = getTile(sheet, 3, 3, TileSize, TileSize)
	Assets["explosion_f5"] = getTile(sheet, 4, 3, TileSize, TileSize)

	Assets["sparkle_f1"] = getTile(sheet, 0, 2, TileSize, TileSize)
	Assets["sparkle_f2"] = getTile(sheet, 1, 2, TileSize, TileSize)
	Assets["sparkle_f3"] = getTile(sheet, 2, 2, TileSize, TileSize)
	Assets["sparkle_f4"] = getTile(sheet, 3, 2, TileSize, TileSize)
	Assets["sparkle_f5"] = getTile(sheet, 4, 2, TileSize, TileSize)
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

func getTile(img *ebiten.Image, x, y, tw, th int) *ebiten.Image {
	return img.SubImage(image.Rect(x*tw, y*th, (x+1)*tw, (y+1)*th)).(*ebiten.Image)
}
