package helper

import (
	"bytes"
	"encoding/xml"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/t-geindre/golem/examples/camera/assets"
	"image"
	"time"
)

type Tileset struct {
	data *TDataSet
	src  *ebiten.Image
	// cache
	imgs  map[int]*ebiten.Image
	tiles map[int]Tile
}

type Tile struct {
	Img       *ebiten.Image
	Animation []TileFrame
}

type TileFrame struct {
	Img      *ebiten.Image
	Duration time.Duration
}

func NewTilesetFromFile(path string) *Tileset {
	fBytes, err := assets.FS.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return NewTilesetFromBytes(fBytes)
}

func NewTilesetFromBytes(src []byte) *Tileset {
	data := &TDataSet{}
	err := xml.Unmarshal(src, data)
	if err != nil {
		panic(err)
	}

	return NewTilesetFromData(data)
}

func NewTilesetFromData(d *TDataSet) *Tileset {
	fBytes := readFile(d.Image.Src)

	src, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(fBytes))
	if err != nil {
		panic(err)
	}

	return &Tileset{
		data:  d,
		src:   src,
		imgs:  make(map[int]*ebiten.Image),
		tiles: make(map[int]Tile),
	}
}

func (t *Tileset) GetTile(id int) Tile {
	if id < 1 {
		panic("id must be greater than 0")
	}
	id--

	tile, ok := t.tiles[id]
	if !ok {
		tile = Tile{
			Img:       t.getImage(id),
			Animation: t.getFrames(id),
		}
		t.tiles[id] = tile
	}

	return tile
}

func (t *Tileset) GetTileSize() (int, int) {
	return t.data.Tw, t.data.Th
}

func (t *Tileset) getImage(id int) *ebiten.Image {
	img, ok := t.imgs[id]

	if !ok {
		x := (id % t.data.Columns) * t.data.Tw
		y := (id / t.data.Columns) * t.data.Th
		img = t.src.SubImage(image.Rect(x, y, x+t.data.Tw, y+t.data.Th)).(*ebiten.Image)
		t.imgs[id] = img
	}

	return img
}

func (t *Tileset) getFrames(id int) []TileFrame {
	for _, tile := range t.data.Tiles {
		if tile.Id == id && len(tile.Frames) > 0 {
			f := make([]TileFrame, 0)
			for _, frame := range tile.Frames {
				f = append(f, TileFrame{
					Img:      t.getImage(frame.TileId),
					Duration: time.Millisecond * time.Duration(frame.Duration),
				})
			}
			return f
		}
	}
	return nil
}
