package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	position := component.GetPosition(e)
	sprite := component.GetSprite(e)

	if position == nil || sprite == nil {
		return
	}

	hw, hh := sprite.Img.Bounds().Dx()/2, sprite.Img.Bounds().Dy()/2

	sprite.Opt.GeoM.Reset()
	sprite.Opt.GeoM.Translate(position.X-float64(hw), position.Y-float64(hh))

	screen.DrawImage(sprite.Img, sprite.Opt)
}
