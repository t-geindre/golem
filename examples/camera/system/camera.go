package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/examples/camera/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type Camera struct {
	zoom float64
	pos  image.Point

	lmx, lmy int
	lm       bool
}

func NewCamera() *Camera {
	return &Camera{
		zoom: 3,
		pos:  image.Point{},
	}
}

func (r *Camera) UpdateOnce(w golem.World) {
	_, y := ebiten.Wheel()
	r.zoom += float64(y) / 10

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		r.lm = false
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		return
	}

	mx, my := ebiten.CursorPosition()
	if r.lm {
		r.pos = r.pos.Add(image.Point{X: mx - r.lmx, Y: my - r.lmy})
	} else {
		r.lm = true
		ebiten.SetCursorShape(ebiten.CursorShapeMove)
	}
	r.lmx, r.lmy = mx, my
}

func (r *Camera) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	sprite := component.GetSprite(e)

	if pos == nil || sprite == nil || r.isCulled(pos, sprite, screen) {
		return
	}

	sh, sw := sprite.Img.Bounds().Dy(), sprite.Img.Bounds().Dx()
	shw, shh := sw/2, sh/2

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(pos.X-shw), float64(pos.Y-shh))
	opts.GeoM.Scale(r.zoom, r.zoom)
	opts.GeoM.Translate(float64(r.pos.X), float64(r.pos.Y))

	screen.DrawImage(sprite.Img, opts)
}

func (r *Camera) isCulled(pos *component.Position, spr *component.Sprite, screen *ebiten.Image) bool {
	// todo panel debug shown/culled
	sprBds := image.Rectangle{Min: image.Point{}, Max: spr.Img.Bounds().Size()}
	sprBds = helper.RectMulF(
		sprBds.Add(pos.Point).Sub(sprBds.Size().Div(2)),
		r.zoom,
	)

	winBds := image.Rectangle{Min: image.Point{}, Max: screen.Bounds().Size()}
	winBds = winBds.Sub(r.pos)

	return !winBds.Overlaps(sprBds)
}
