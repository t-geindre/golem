package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/nodes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
)

type Bounce struct {
	bds image.Rectangle
}

func NewBounce() *Bounce {
	return &Bounce{}
}

func (b *Bounce) UpdateOnce(w golem.World, c golem.Clock) {
	ww, wh := ebiten.WindowSize()
	b.bds = image.Rect(0, 0, ww, wh)
}

func (b *Bounce) Update(e golem.Entity, w golem.World, _ golem.Clock) {
	eBds := component.GetBoundaries(e)
	eVel := component.GetVelocity(e)

	if eBds == nil || eVel == nil {
		return
	}

	bds := b.bds
	rBds := eBds.Rectangle
	if p := w.GetParentEntity(); p != nil {
		if pb := component.GetBoundaries(p); pb != nil {
			bds = pb.Rectangle
			rBds = eBds.Add(pb.Rectangle.Min)
		}
	}

	if rBds.Min.X < bds.Min.X {
		eBds.Rectangle = eBds.Rectangle.Add(image.Pt(bds.Min.X-rBds.Min.X, 0))
		eVel.X = -eVel.X
	}
	if rBds.Min.Y < bds.Min.Y {
		eBds.Rectangle = eBds.Rectangle.Add(image.Pt(0, bds.Min.Y-rBds.Min.Y))
		eVel.Y = -eVel.Y
	}
	if rBds.Max.X > bds.Max.X {
		eBds.Rectangle = eBds.Rectangle.Add(image.Pt(bds.Max.X-rBds.Max.X, 0))
		eVel.X = -eVel.X
	}
	if rBds.Max.Y > bds.Max.Y {
		eBds.Rectangle = eBds.Rectangle.Add(image.Pt(0, bds.Max.Y-rBds.Max.Y))
		eVel.Y = -eVel.Y
	}
}
