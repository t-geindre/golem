package component

import (
	"image"
)

//go:generate golem component Camera
type Camera struct {
	Projection   image.Rectangle
	ProjIsScreen bool

	// Field of view
	// computed according to the projection size and the zoom
	Fov image.Rectangle

	// Zoom control
	Zoom, TargetZoom    float64
	ZoomSpeed, ZoomFact float64
	ZoomCap             image.Point

	// Controls
	WheelZoom    bool
	KeyMove      bool
	KeyMoveSpeed int
	MouseMove    bool
}

func NewCamera(z, tz float64) *Camera {
	return &Camera{
		ProjIsScreen: true,
		Zoom:         z,
		TargetZoom:   tz,
		ZoomSpeed:    .3,
		ZoomFact:     .6,
		ZoomCap:      image.Pt(1, 10),

		WheelZoom:    true,
		KeyMove:      true,
		KeyMoveSpeed: 4,
		MouseMove:    true,
	}
}
