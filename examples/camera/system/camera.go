package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/examples/camera/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
	"math"
)

type Camera struct {
	zoom, tZoom float64
	pos         image.Point

	lmx, lmy int
	lm       bool
}

func NewCamera() *Camera {
	return &Camera{
		zoom:  3,
		tZoom: 3,
		pos:   image.Point{},
	}
}

func (r *Camera) UpdateOnce(w golem.World) {
	_, y := ebiten.Wheel()
	if y != 0 {
		delta := float64(y) / 5
		r.tZoom += delta
	}

	if r.zoom != r.tZoom {
		d := (r.tZoom - r.zoom) / 3
		if math.Abs(d) < 0.01 {
			r.zoom = r.tZoom
		} else {
			r.zoom += d
		}
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		r.lm = false
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	} else {
		mx, my := ebiten.CursorPosition()
		if r.lm {
			r.pos = r.pos.Add(image.Point{X: mx - r.lmx, Y: my - r.lmy})
		} else {
			r.lm = true
			ebiten.SetCursorShape(ebiten.CursorShapeMove)
		}
		r.lmx, r.lmy = mx, my
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		r.pos.Y += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		r.pos.Y -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		r.pos.X += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		r.pos.X -= 5
	}
}

func (r *Camera) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	pos := component.GetPosition(e)
	sprite := component.GetSprite(e)

	if pos == nil || sprite == nil || r.isCulled(pos, sprite, screen) {
		return
	}

	shw, shh := sprite.Img.Bounds().Dy()/2, sprite.Img.Bounds().Dx()/2

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
