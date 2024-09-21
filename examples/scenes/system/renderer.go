package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"math"
	_ "unsafe"
)

type Renderer struct {
	ww, wh float64
	dirty  bool
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) UpdateOnce(w golem.World) {
	ww, wh := ebiten.Monitor().Size()
	//if ebiten.IsFullscreen() {
	//	ww, wh = ebiten.Monitor().Size()
	//} else {
	//	ww, wh = ebiten.WindowSize()
	//}

	wwf, whf := float64(ww), float64(wh)
	if r.ww != wwf || r.wh != whf {
		r.ww, r.wh = wwf, whf
		r.dirty = true
		return
	}

	r.dirty = false
}

func (r *Renderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	if pos == nil {
		return
	}

	spr := component.GetSprite(e)
	txt := component.GetText(e)

	if spr == nil && txt == nil {
		return
	}

	opts := &ebiten.DrawImageOptions{}
	r.applyOpts(e, opts)
	r.applyOpts(w.GetParentEntity(), opts)
	opts.GeoM.Translate(pos.RelX*r.ww, pos.RelY*r.wh)

	if spr != nil {
		r.renderSprite(spr, pos, opts, screen)
	}

	if txt != nil {
		r.renderText(e, txt, pos, opts, screen)
	}

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

func (r *Renderer) renderSprite(spr *component.Sprite, pos *component.Position, opts *ebiten.DrawImageOptions, screen *ebiten.Image) {
	sw, sh := spr.Img.Bounds().Dx(), spr.Img.Bounds().Dy()
	swf, shf := float64(sw), float64(sh)
	/*
		scale := math.Min(r.ww, r.wh) / 8
		sx, sy := scale/swf, scale/shf
		opts.GeoM.Scale(sx, sy)
	*/
	opts.GeoM.Translate(-swf*pos.OrigX, -shf*pos.OrigY)

	screen.DrawImage(spr.Img, opts)
}

func (r *Renderer) renderText(e golem.Entity, txt *component.Text, pos *component.Position, opts *ebiten.DrawImageOptions, screen *ebiten.Image) {
	r.applyTextColor(e, opts)
	r.computeTextFace(txt)
	r.computeTextSize(txt)
	opts.Filter = ebiten.FilterLinear // Improve text quality

	opts.GeoM.Translate(-txt.Width*pos.OrigX, txt.FontAscent-txt.Height*pos.OrigY)
	text.Draw(screen, txt.GetValue(), txt.FontFace, &text.DrawOptions{DrawImageOptions: *opts})
}

func (r *Renderer) applyTextColor(e golem.Entity, opts *ebiten.DrawImageOptions) {
	col := colornames.Black
	clr := component.GetColor(e)
	if clr != nil {
		col = clr.Value
	}
	opts.ColorScale.ScaleWithColor(col)
}

func (r *Renderer) computeTextFace(txt *component.Text) {
	if txt.FontFace != nil && !r.dirty {
		return
	}

	fontSize := txt.FontSize * math.Min(r.ww, r.wh) / 100

	face, _ := opentype.NewFace(txt.Font, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	txt.FontFace = text.NewGoXFace(face)
	txt.FontAscent = txt.FontFace.Metrics().VAscent
	txt.Dirty = true
}

func (r *Renderer) computeTextSize(txt *component.Text) {
	if txt.Width != 0 && txt.Height != 0 && !txt.Dirty {
		return
	}

	txt.Width, txt.Height = text.Measure(txt.GetValue(), txt.FontFace, 0)
}
