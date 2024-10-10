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
	lmx, lmy int
	lCap     bool
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) Update(e golem.Entity, _ golem.World, _ golem.Clock) {
	cam := component.GetCamera(e)
	if cam == nil {
		return
	}

	c.wheelZoom(cam)
	c.zoomUpdate(cam)
	c.projectionUpdate(cam)

	pos := component.GetPosition(e)
	if pos == nil {
		return
	}

	c.fovUpdate(cam, pos)
	c.keyMove(cam, pos)
	c.mouseMove(cam, pos)
}
func (c *Camera) projectionUpdate(cam *component.Camera) {
	if cam.ProjIsScreen {
		cam.Projection.Min.X, cam.Projection.Min.Y = 0, 0
		cam.Projection.Max.X, cam.Projection.Max.Y = ebiten.WindowSize()
	}
}

func (c *Camera) fovUpdate(cam *component.Camera, pos *component.Position) {
	cam.Fov = image.Rect(
		-cam.Projection.Dx()/2, -cam.Projection.Dy()/2,
		cam.Projection.Dx()/2, cam.Projection.Dy()/2,
	)
	cam.Fov = helper.RectMulF(cam.Fov, 1/cam.Zoom)

	cam.Fov.Min = cam.Fov.Min.Add(pos.Point)
	cam.Fov.Max = cam.Fov.Max.Add(pos.Point)
}

func (c *Camera) zoomUpdate(cam *component.Camera) {
	if cam.TargetZoom != cam.Zoom {
		d := (cam.TargetZoom - cam.Zoom) * cam.ZoomSpeed
		if math.Abs(d) < 0.01 {
			cam.Zoom = cam.TargetZoom
		} else {
			cam.Zoom += d
		}
	}
}

func (c *Camera) wheelZoom(cam *component.Camera) {
	if !cam.WheelZoom {
		return
	}

	_, y := ebiten.Wheel()
	if y != 0 {
		cam.TargetZoom += y * cam.ZoomFact
		if cam.TargetZoom < float64(cam.ZoomCap.X) {
			cam.TargetZoom = float64(cam.ZoomCap.X)
		}
		if cam.TargetZoom > float64(cam.ZoomCap.Y) {
			cam.TargetZoom = float64(cam.ZoomCap.Y)
		}
	}
}

func (c *Camera) keyMove(cam *component.Camera, pos *component.Position) {
	if !cam.KeyMove {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		pos.Y -= cam.KeyMoveSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		pos.Y += cam.KeyMoveSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		pos.X -= cam.KeyMoveSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		pos.X += cam.KeyMoveSpeed
	}
}

func (c *Camera) mouseMove(cam *component.Camera, pos *component.Position) {
	if !cam.MouseMove {
		return
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		c.lCap = false
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	} else {
		mx, my := ebiten.CursorPosition()
		if c.lCap {
			dx := int(math.Round(float64(mx-c.lmx) / cam.Zoom))
			dy := int(math.Round(float64(my-c.lmy) / cam.Zoom))
			pos.Point = pos.Sub(image.Pt(dx, dy))
		} else {
			c.lCap = true
			ebiten.SetCursorShape(ebiten.CursorShapeMove)
		}
		c.lmx, c.lmy = mx, my
	}
}
