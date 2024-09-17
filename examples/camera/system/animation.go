package system

import (
	"github.com/t-geindre/golem/examples/camera/component"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Animation struct {
}

func NewAnimation() *Animation {
	return &Animation{}
}

func (a *Animation) Update(e golem.Entity, w golem.World) {
	anim := component.GetAnimation(e)
	sprite := component.GetSprite(e)

	if anim == nil || sprite == nil || len(anim.Frames) == 0 {
		return
	}

	if anim.Start.IsZero() {
		anim.Start = time.Now()
		anim.Current = 0
		sprite.Img = anim.Frames[0].Img
		return
	}

	if time.Since(anim.Start) > anim.Frames[anim.Current].Duration {
		anim.Current = (anim.Current + 1) % len(anim.Frames)
		sprite.Img = anim.Frames[anim.Current].Img
		anim.Start = time.Now()
	}
}
