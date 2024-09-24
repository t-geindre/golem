package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
	"math"
)

type Renderer struct {
	srw, srh float64
	ww, wh   float64
	scale    float64
}

func NewRenderer(srw, srh float64) *Renderer {
	return &Renderer{
		srw: srw,
		srh: srh,
	}
}

func (r *Renderer) UpdateOnce(w golem.World) {
	ww, wh := ebiten.Monitor().Size()
	if !ebiten.IsFullscreen() {
		ww, wh = ebiten.WindowSize()
	}

	wwf, whf := float64(ww), float64(wh)
	r.ww, r.wh = wwf, whf

	r.scale = math.Min(wwf/r.srw, whf/r.srh)
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
	if bds != nil && bds.StickScreen {
		bds.Rectangle = image.Rect(0, 0, int(r.ww), int(r.wh))
	}

	rot := component.GetRotation(e)
	if rot != nil && bds != nil && rot.Angle != 0 {
		opts.GeoM.Translate(-float64(bds.Dx())/2, -float64(bds.Dy())/2)
		opts.GeoM.Rotate(rot.Angle)
		opts.GeoM.Translate(float64(bds.Dx())/2, float64(bds.Dy())/2)
		opts.Filter = ebiten.FilterLinear // Improve rendering quality
	}

	op := component.GetOpacity(e)
	if op != nil {
		opts.ColorScale.ScaleAlpha(op.Value)
	}

	scl := component.GetScale(e)
	if scl != nil && bds != nil {
		v := scl.Value
		scrScale := component.GetScreenScale(e)
		if scrScale != nil {
			v = v * r.scale * scrScale.Value
		}
		if v < 0 {
			v = 0
		}
		opts.GeoM.Scale(v, v)
		opts.GeoM.Translate(
			float64(bds.Dx())*scl.OriginX*(1-v),
			float64(bds.Dy())*scl.OriginX*(1-v),
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
