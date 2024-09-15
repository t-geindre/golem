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

	plSheet := getAssetImage("player_16_32.png")
	Assets["player_f1"] = getTile(plSheet, 2, 0, TileSize, TileSize*2)
	Assets["player_f2"] = getTile(plSheet, 3, 0, TileSize, TileSize*2)

	enLipsSheet := getAssetImage("en_lips_16.png")
	Assets["en_lips_f1"] = getTile(enLipsSheet, 0, 0, TileSize, TileSize)
	Assets["en_lips_f2"] = getTile(enLipsSheet, 1, 0, TileSize, TileSize)
	Assets["en_lips_f3"] = getTile(enLipsSheet, 2, 0, TileSize, TileSize)
	Assets["en_lips_f4"] = getTile(enLipsSheet, 3, 0, TileSize, TileSize)
	Assets["en_lips_f5"] = getTile(enLipsSheet, 4, 0, TileSize, TileSize)

	enAllanSheet := getAssetImage("en_allan_16.png")
	Assets["en_allan_f1"] = getTile(enAllanSheet, 0, 0, TileSize, TileSize)
	Assets["en_allan_f2"] = getTile(enAllanSheet, 1, 0, TileSize, TileSize)
	Assets["en_allan_f3"] = getTile(enAllanSheet, 2, 0, TileSize, TileSize)
	Assets["en_allan_f4"] = getTile(enAllanSheet, 3, 0, TileSize, TileSize)
	Assets["en_allan_f5"] = getTile(enAllanSheet, 4, 0, TileSize, TileSize)
	Assets["en_allan_f6"] = getTile(enAllanSheet, 4, 0, TileSize, TileSize)

	enBonbonSheet := getAssetImage("en_bonbon_16.png")
	Assets["en_bonbon_f1"] = getTile(enBonbonSheet, 0, 0, TileSize, TileSize)
	Assets["en_bonbon_f2"] = getTile(enBonbonSheet, 1, 0, TileSize, TileSize)
	Assets["en_bonbon_f3"] = getTile(enBonbonSheet, 2, 0, TileSize, TileSize)
	Assets["en_bonbon_f4"] = getTile(enBonbonSheet, 3, 0, TileSize, TileSize)

	bulletSheet := getAssetImage("bullet_16.png")
	Assets["bullet_f1"] = getTile(bulletSheet, 0, 0, TileSize, TileSize)
	Assets["bullet_f2"] = getTile(bulletSheet, 1, 0, TileSize, TileSize)

	expSheet := getAssetImage("explosion_16.png")
	Assets["explosion_f1"] = getTile(expSheet, 0, 0, TileSize, TileSize)
	Assets["explosion_f2"] = getTile(expSheet, 1, 0, TileSize, TileSize)
	Assets["explosion_f3"] = getTile(expSheet, 2, 0, TileSize, TileSize)
	Assets["explosion_f4"] = getTile(expSheet, 3, 0, TileSize, TileSize)
	Assets["explosion_f5"] = getTile(expSheet, 4, 0, TileSize, TileSize)

	spkSheet := getAssetImage("sparkle_16.png")
	Assets["sparkle_f1"] = getTile(spkSheet, 0, 0, TileSize, TileSize)
	Assets["sparkle_f2"] = getTile(spkSheet, 1, 0, TileSize, TileSize)
	Assets["sparkle_f3"] = getTile(spkSheet, 2, 0, TileSize, TileSize)
	Assets["sparkle_f4"] = getTile(spkSheet, 3, 0, TileSize, TileSize)
	Assets["sparkle_f5"] = getTile(spkSheet, 4, 0, TileSize, TileSize)
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

func getTile(img *ebiten.Image, x, y, tw, th int) *ebiten.Image {
	return img.SubImage(image.Rect(x*tw, y*th, (x+1)*tw, (y+1)*th)).(*ebiten.Image)
}
