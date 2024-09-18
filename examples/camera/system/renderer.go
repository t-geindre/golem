package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/examples/camera/entity"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type Renderer struct {
	cam *entity.Camera
}

func NewRenderer(cam *entity.Camera) *Renderer {
	return &Renderer{
		cam: cam,
	}
}

func (r *Renderer) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	sprite := component.GetSprite(e)

	if pos == nil || sprite == nil {
		return
	}

	shw, shh := sprite.Img.Bounds().Dy()/2, sprite.Img.Bounds().Dx()/2

	// Culling
	sprBounds := sprite.Img.Bounds()
	sprBounds = sprBounds.Sub(sprBounds.Min).Add(pos.Point).Sub(image.Pt(shw, shh))
	if !sprBounds.Overlaps(r.cam.Fov) {
		return
	}

	// Rendering
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(pos.X-shw), float64(pos.Y-shh))
	opts.GeoM.Translate(float64(-r.cam.Fov.Min.X), float64(-r.cam.Fov.Min.Y))
	opts.GeoM.Scale(
		float64(r.cam.Projection.Dx())/float64(r.cam.Fov.Dx()),
		float64(r.cam.Projection.Dy())/float64(r.cam.Fov.Dy()),
	)

	screen.DrawImage(sprite.Img, opts)
}
