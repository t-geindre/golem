package entity

import (
	"github.com/t-geindre/golem/examples/nodes/component"
	"github.com/t-geindre/golem/examples/nodes/helper"
	"github.com/t-geindre/golem/pkg/golem"
)

type Node struct {
	golem.Entity
	golem.World
	*component.Boundaries
	*component.Velocity
	*component.Geometry
	*component.Color
}

func NewNode(l golem.LayerID, vx, vy, x, y, w, h int) *Node {
	n := &Node{
		Entity:     golem.NewEntity(l),
		World:      golem.NewWorld(),
		Boundaries: component.NewBoundaries(x, y, w, h),
		Velocity:   component.NewVelocity(vx, vy),
		Geometry:   component.NewGeometry(),
		Color:      component.NewColor(helper.RandomColor(0xcc), helper.RandomColor(0xff)),
	}

	n.SetParentSharedSystems(true)

	return n
}
