package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Renderer struct {
	parent golem.Entity
}

func NewRenderer(parent golem.Entity) *Renderer {
	return &Renderer{
		parent: parent,
	}
}

func (r *Renderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	sprite := component.GetSprite(e)

	if pos == nil || sprite == nil {
		return
	}

	hw, hh := sprite.Img.Bounds().Dx()/2, sprite.Img.Bounds().Dy()/2

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(pos.X-float64(hw), pos.Y-float64(hh))

	if r.parent != nil {
		parentPos := component.GetPosition(r.parent)
		if parentPos != nil {
			opt.GeoM.Translate(parentPos.X, parentPos.Y)
		}
	}

	screen.DrawImage(sprite.Img, opt)
}
