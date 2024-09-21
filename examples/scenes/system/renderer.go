package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type Renderer struct {
	ww, wh float64
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) UpdateOnce(w golem.World) {
	ww, wh := ebiten.Monitor().Size()
	wwf, whf := float64(ww), float64(wh)
	r.ww, r.wh = wwf, whf
}

func (r *Renderer) getDrawOpts(
	e golem.Entity,
	w golem.World,
	pos *component.Position,
	bds image.Rectangle,
) *ebiten.DrawImageOptions {
	opts := &ebiten.DrawImageOptions{}

	r.applyOpts(e, opts, bds)
	r.applyOpts(w.GetParentEntity(), opts, bds)

	hw, hh := float64(bds.Dx())*pos.OrigX, float64(bds.Dy())*pos.OrigY
	opts.GeoM.Translate(pos.RelX*r.ww-hw, pos.RelY*r.wh-hh)

	return opts
}

func (r *Renderer) applyOpts(e golem.Entity, opts *ebiten.DrawImageOptions, bds image.Rectangle) {
	if e == nil {
		return
	}

	opacity := component.GetOpacity(e)
	if opacity != nil {
		opts.ColorScale.ScaleAlpha(opacity.Value)
	}

	scale := component.GetScale(e)
	if scale != nil {
		opts.GeoM.Scale(scale.Value, scale.Value)
		opts.GeoM.Translate(
			float64(bds.Dx())/2*(1-scale.Value),
			float64(bds.Dy())/2*(1-scale.Value),
		)
	}

	color := component.GetColor(e)
	if color != nil {
		opts.ColorScale.ScaleWithColor(color.Value)
	}
}
