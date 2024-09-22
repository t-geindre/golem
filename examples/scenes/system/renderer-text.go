package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type TextRenderer struct {
	*Renderer
}

func NewTextRenderer() *TextRenderer {
	return &TextRenderer{
		Renderer: NewRenderer(),
	}
}

func (r *TextRenderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	txt := component.GetText(e)

	if pos == nil || txt == nil {
		return
	}

	r.computeBounds(txt)

	opts := &text.DrawOptions{
		DrawImageOptions: *r.getDrawOpts(e, w, pos, txt.Bounds),
		LayoutOptions: text.LayoutOptions{
			LineSpacing: txt.LineHeight,
		},
	}

	text.Draw(screen, txt.Text, txt.Face, opts)
}

func (r *TextRenderer) computeBounds(txt *component.Text) {
	if !txt.Bounds.Empty() {
		return
	}

	if len(txt.Text) == 0 {
		return
	}

	m := txt.Face.Metrics()
	txt.LineHeight = m.HLineGap + m.HAscent + m.HDescent

	w, h := text.Measure(txt.Text, txt.Face, txt.LineHeight)
	txt.Bounds = image.Rect(0, 0, int(w), int(h))

}
