package helper

import (
	"encoding/json"
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/examples/camera/entity"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type Map struct {
	data    *MDataMap
	tileset *Tileset
}

func NewMapFromFile(path string) *Map {
	return NewMapFromBytes(readFile(path))
}

func NewMapFromBytes(scr []byte) *Map {
	m := &MDataMap{}
	err := json.Unmarshal(scr, &m)

	if err != nil {
		panic(err)
	}

	return NewMapFromData(m)
}

func NewMapFromData(data *MDataMap) *Map {
	return &Map{
		data: data,
	}
}

func (m *Map) Size() image.Point {
	return image.Pt(m.data.Width*m.data.Tw, m.data.Height*m.data.Th)
}

func (m *Map) Center() image.Point {
	return m.Size().Div(2)
}

func (m *Map) Layers() []golem.LayerID {
	ls := make([]golem.LayerID, len(m.data.Layers))
	for i := range m.data.Layers {
		ls[i] = golem.LayerID(i)
	}
	return ls
}

func (m *Map) Entities() []golem.Entity {
	entities := make([]golem.Entity, 0)
	m.loadTileset()

	for lid, layer := range m.data.Layers {
		for idx, tid := range layer.Data {
			if tid == 0 {
				continue
			}
			entities = append(entities, m.getTileEntity(lid, tid, idx))
		}
	}

	return entities
}

func (m *Map) loadTileset() {
	if m.tileset != nil {
		return
	}

	if len(m.data.Tilesets) == 0 {
		panic("no tileset found in map")
	}
	if len(m.data.Tilesets) > 1 {
		panic("multiple tilesets not supported")
	}

	m.tileset = NewTilesetFromFile(m.data.Tilesets[0].Source)
}

func (m *Map) getTileEntity(lid, tid, idx int) golem.Entity {
	pos := image.Pt(
		idx%m.data.Width*m.data.Tw,
		idx/m.data.Height*m.data.Th,
	)

	tile := m.tileset.GetTile(tid)

	fs := make([]component.Frame, 0)
	if tile.Animation != nil {
		for _, f := range tile.Animation {
			fs = append(fs, component.Frame{Img: f.Img, Duration: f.Duration})
		}
	} else {
		fs = append(fs, component.Frame{Img: tile.Img})
	}

	return entity.NewTile(golem.LayerID(lid), pos, fs...)
}
