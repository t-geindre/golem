package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type SpriteRenderer struct {
	*Renderer
}

func NewSpriteRenderer() *SpriteRenderer {
	return &SpriteRenderer{
		Renderer: NewRenderer(),
	}
}

func (r *SpriteRenderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	spr := component.GetSprite(e)

	if pos == nil || spr == nil {
		return
	}

	opts := r.getDrawOpts(e, w)

	screen.DrawImage(spr.Img, opts)
}
