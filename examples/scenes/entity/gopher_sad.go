package entity

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type GopherSad struct {
	golem.Entity
	*component.Position
	*component.Sprite
	*component.Animation
}

func NewGopherSad(l golem.LayerID, x, y float64) *Gopher {
	return &Gopher{
		Entity:   golem.NewEntity(l),
		Position: component.NewPosition(x, y),
		Sprite:   component.NewSprite(helper.Assets[0]),
		Animation: component.NewAnimation(
			component.NewFrame(helper.Assets[22], time.Second*3),
			component.NewFrame(helper.Assets[13], time.Second*1),
			component.NewFrame(helper.Assets[18], time.Second*1),
			component.NewFrame(helper.Assets[13], time.Second*1),
		),
	}
}
