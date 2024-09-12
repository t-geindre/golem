package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type GopherAngry struct {
	golem.Entity
	*component.Position
	*component.Sprite
	*component.Animation
}

func NewGopherAngry(l golem.LayerID, x, y float64) *Gopher {
	return &Gopher{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(x, y),
		Sprite:   component.NewSprite(helper.Assets[0]),
		Animation: component.NewAnimation(
			component.NewFrame(helper.Assets[15], time.Second*3),
			component.NewFrame(helper.Assets[20], time.Second*3),
			component.NewFrame(helper.Assets[1], time.Millisecond*200),
			component.NewFrame(helper.Assets[20], time.Second*3),
			component.NewFrame(helper.Assets[1], time.Millisecond*200),
		),
	}
}
