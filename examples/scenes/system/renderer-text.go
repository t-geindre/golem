package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
)

type TextRenderer struct {
	*Renderer
}

func NewTextRenderer(srw, srh float64) *TextRenderer {
	return &TextRenderer{
		Renderer: NewRenderer(srw, srh),
	}
}

func (r *TextRenderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	txt := component.GetText(e)
	bds := component.GetBoundaries(e)

	if pos == nil || txt == nil || bds == nil {
		return
	}

	r.computeFace(txt, bds)
	imgOpts := r.getDrawOpts(e, w)

	opts := &text.DrawOptions{
		DrawImageOptions: *imgOpts,
		LayoutOptions: text.LayoutOptions{
			LineSpacing: txt.LineHeight,
		},
	}

	text.Draw(screen, txt.Text, txt.Face, opts)
}

func (r *TextRenderer) computeFace(txt *component.Text, bds *component.Boundaries) {
	if txt.Face != nil && txt.Scale == r.scale {
		return
	}

	face, err := opentype.NewFace(txt.Font, &opentype.FaceOptions{
		Size:    txt.Size * r.scale,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	txt.Face = text.NewGoXFace(face)
	txt.Scale = r.scale

	r.computeBounds(txt, bds)
}

func (r *TextRenderer) computeBounds(txt *component.Text, bds *component.Boundaries) {
	m := txt.Face.Metrics()
	txt.LineHeight = m.HLineGap + m.HAscent + m.HDescent

	w, h := text.Measure(txt.Text, txt.Face, txt.LineHeight)
	txt.Bounds = image.Rect(0, 0, int(w), int(h))

	bds.Rectangle = image.Rect(0, 0, int(w), int(h))
}
