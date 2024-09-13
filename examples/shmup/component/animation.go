package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Frame struct {
	Img      *ebiten.Image
	Duration time.Duration
}

func NewFrame(img *ebiten.Image, duration time.Duration) Frame {
	return Frame{
		Img:      img,
		Duration: duration,
	}
}

//go:generate golem component Animation
type Animation struct {
	Frames  []Frame
	Current int
	Start   time.Time
}

func NewAnimation(frames ...Frame) *Animation {
	return &Animation{
		Frames: frames,
	}
}
