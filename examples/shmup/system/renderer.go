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

	opt := &ebiten.DrawImageOptions{}

	rot := component.GetRotation(e)
	if rot != nil {
		opt.GeoM.Translate(float64(-hw), float64(-hh))
		opt.GeoM.Rotate(rot.Angle)
		opt.GeoM.Translate(float64(hw), float64(hh))
	}

	opt.GeoM.Translate(position.X-float64(hw), position.Y-float64(hh))

	screen.DrawImage(sprite.Img, opt)
}
