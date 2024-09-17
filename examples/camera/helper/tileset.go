package helper

import (
	"encoding/xml"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func NewTileset(src []byte) *Tileset {
	data := &TDataSet{}
	err := xml.Unmarshal(src, data)
	if err != nil {
		panic(err)
	}

	return &Tileset{
		src:     src,
		Columns: columns,
		cache:   make(map[int]*ebiten.Image),
		Tw:      tw,
		Th:      th,
	}
}

func (t *Tileset) GetTile(index int) *ebiten.Image {
	if index < 1 {
		panic("index must be greater than 0")
	}

	index--

	if img, ok := t.cache[index]; ok {
		return img
	}

	x := (index % t.Columns) * t.Tw
	y := (index / t.Columns) * t.Th
	img := t.src.SubImage(image.Rect(x, y, x+t.Tw, y+t.Th)).(*ebiten.Image)
	t.cache[index] = img

	return img
}
