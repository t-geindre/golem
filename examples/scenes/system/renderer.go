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
	ww, wh := ebiten.WindowSize()
	r.ww, r.wh = float64(ww), float64(wh)
}

func (r *Renderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	sprite := component.GetSprite(e)

	if pos == nil || sprite == nil {
		return
	}

	hw, hh := sprite.Img.Bounds().Dx()/2, sprite.Img.Bounds().Dy()/2

	opts := &ebiten.DrawImageOptions{}

	r.applyOpts(e, opts)
	r.applyOpts(w.GetParentEntity(), opts)

	opts.GeoM.Translate(pos.RelX*r.ww-float64(hw), pos.RelY*r.wh-float64(hh))

	screen.DrawImage(sprite.Img, opts)
}

func (r *Renderer) applyOpts(e golem.Entity, opts *ebiten.DrawImageOptions) {
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
	}
}
