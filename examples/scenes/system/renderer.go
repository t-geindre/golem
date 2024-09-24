package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
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
) *ebiten.DrawImageOptions {
	opts := &ebiten.DrawImageOptions{}

	r.applyOpts(e, opts)
	r.applyOpts(w.GetParentEntity(), opts)

	return opts
}

func (r *Renderer) applyOpts(e golem.Entity, opts *ebiten.DrawImageOptions) {
	if e == nil {
		return
	}

	bds := component.GetBoundaries(e)

	opacity := component.GetOpacity(e)
	if opacity != nil {
		opts.ColorScale.ScaleAlpha(opacity.Value)
	}

	scale := component.GetScale(e)
	if scale != nil && bds != nil {
		v := scale.Value
		if v < 0 {
			v = 0
		}
		opts.GeoM.Scale(v, v)
		opts.GeoM.Translate(
			float64(bds.Dx())*scale.OriginX*(1-v),
			float64(bds.Dy())*scale.OriginX*(1-v),
		)
	}

	color := component.GetColor(e)
	if color != nil {
		opts.ColorScale.ScaleWithColor(color.Value)
	}

	pos := component.GetPosition(e)
	if pos != nil && bds != nil {
		hw, hh := float64(bds.Dx())*pos.OriginX, float64(bds.Dy())*pos.OriginY
		opts.GeoM.Translate(pos.RelX*r.ww-hw, pos.RelY*r.wh-hh)
	}
}
