package system

import (
	"github.com/t-geindre/golem/examples/shmup/component"
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

	if anim == nil || sprite == nil {
		return
	}

	if anim.Start.IsZero() {
		anim.Start = time.Now()
		anim.Current = 0
		sprite.Img = anim.Frames[0].Img
		return
	}

	if time.Since(anim.Start) > anim.Frames[anim.Current].Duration {
		if anim.Current == len(anim.Frames)-1 {
			if !anim.Loop {
				return
			}
		}
		anim.Current = (anim.Current + 1) % len(anim.Frames)
		sprite.Img = anim.Frames[anim.Current].Img
		anim.Start = time.Now()
	}
}
