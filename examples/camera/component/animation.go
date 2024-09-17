package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

//go:generate golem component Animation
type Animation struct {
	Frames  []Frame
	Current int
	Start   time.Time
	Loop    bool
}

func NewAnimation(loop bool, frames ...Frame) *Animation {
	return &Animation{
		Frames: frames,
		Loop:   loop,
	}
}

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
