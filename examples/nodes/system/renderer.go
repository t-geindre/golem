package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/t-geindre/golem/examples/nodes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (d *Renderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	geo := component.GetGeometry(e)
	if geo == nil {
		return
	}

	bds := component.GetBoundaries(e)
	if bds == nil {
		return
	}

	col := component.GetColor(e)
	if col == nil {
		return
	}

	xs, ys := geo.GeoM.Apply(float64(0), float64(0))
	xe, ye := geo.GeoM.Apply(float64(bds.Rectangle.Dx()), float64(bds.Rectangle.Dy()))

	vector.DrawFilledRect(
		screen,
		float32(xs), float32(ys),
		float32(xe-xs), float32(ye-ys),
		col.Background,
		false,
	)

	vector.StrokeRect(
		screen,
		float32(xs), float32(ys),
		float32(xe-xs), float32(ye-ys),
		2, col.Borders,
		false,
	)
}
