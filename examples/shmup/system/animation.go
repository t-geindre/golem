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

func (a *Animation) Update(e golem.Entity, _ golem.World, _ golem.Clock) {
	anim := component.GetAnimation(e)
	sprite := component.GetSprite(e)

	if anim == nil || sprite == nil {
		return
	}

	set := component.GetAnimationSet(e)
	if set != nil && set.Next != set.Current {
		if _, ok := set.Animations[set.Next]; !ok {
			set.Next = set.Default
		}

		anim.Frames = set.Animations[set.Next].Frames
		anim.Current = 0
		set.Current = set.Next
		set.Current = set.Next
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
